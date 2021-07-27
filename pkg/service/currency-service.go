package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/bino1490/crypto-svc/pkg/entity"
	"github.com/bino1490/crypto-svc/pkg/logger"
)

//regular expression for the get. pose emthods
var (
	getRe      = regexp.MustCompile(`^\/currency\/all[\/]*$`)
	getByKeyRe = regexp.MustCompile(`^\/currency\/[^*]`)
	createRe   = regexp.MustCompile(`^\/currency[\/]*$`)
)

//creating interface with abstract methods
type MemoryService interface {
	InMemGetPOST(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetDataByKey(w http.ResponseWriter, r *http.Request)
}

type MemHandlers struct {
	sync.Mutex
}

func NewInMemService() *MemHandlers {
	return &MemHandlers{}

}

//InMemGetPOST to perform in memory Add/GET operations
func (h *MemHandlers) InMemGetPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && getRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	// case r.Method == http.MethodPost && createRe.MatchString(r.URL.Path):
	// 	h.Post(w, r)
	// 	return
	case r.Method == http.MethodGet && getByKeyRe.MatchString(r.URL.Path):
		h.GetDataByKey(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.BootstrapLogger.Debug("Method Not Allowd")
		w.Write([]byte("Method not allowed"))
		return
	}
}

// Get all the data from inmemory
func (h *MemHandlers) Get(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug("Entering Service.Get() ...")

	err, bodyBytes := GetRequest("https://api.hitbtc.com/api/2/public/ticker")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to Fetch Data From api.hitbtc"))
	}
	var FinalResult []entity.FinalRequest
	var responseObject []entity.FinalRequest
	json.Unmarshal(bodyBytes, &responseObject)

	err, sym := GetRequest("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil || len(responseObject) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to Fetch Data From api.hitbtc"))
	}
	var symObject []entity.SymbolRequest
	json.Unmarshal(sym, &symObject)

	err, cur := GetRequest("https://api.hitbtc.com/api/2/public/currency")
	if err != nil || len(symObject) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to Fetch Data From api.hitbtc"))
	}
	var curObject []entity.CurrencyRequest
	json.Unmarshal(cur, &curObject)

	for _, resp := range responseObject {
		if resp.Symbol != "" {
			for _, sp := range symObject {
				if resp.Symbol == sp.ID {
					resp.FeeCurrency = sp.FeeCurrency
					resp.ID = sp.BaseCurrency
					for _, cr := range curObject {
						if cr.ID == sp.BaseCurrency {
							resp.FullName = cr.FullName
						}

					}
				}
			}

		}
		resp.Symbol = ""
		FinalResult = append(FinalResult, resp)

	}
	logger.Logger.Debug("API Response as struct ", FinalResult)
	response := entity.ResponseData{
		Currencies: FinalResult,
	}

	res, err := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Get specific Key from Inmemory
func (h *MemHandlers) GetDataByKey(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug("Entering service.GetDataByKey() ...")
	matches := strings.Split(r.URL.Path, "/")

	if len(matches) <= 2 {
		notFound(w, r)
		return
	}
	h.Lock()
	reqparam := matches[2]
	err, bodyBytes := GetRequest("https://api.hitbtc.com/api/2/public/ticker/" + reqparam)
	var responseObject entity.FinalRequest
	json.Unmarshal(bodyBytes, &responseObject)
	if err != nil || responseObject.Symbol == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to Fetch Data From api.hitbtc Ticker"))
	}
	symerr, sym := GetRequest("https://api.hitbtc.com/api/2/public/symbol/" + responseObject.Symbol)
	var symObject entity.SymbolRequest
	json.Unmarshal(sym, &symObject)
	if symerr != nil || symObject.BaseCurrency == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to Fetch Data From api.hitbtc Symbol"))
	}
	responseObject.FeeCurrency = symObject.FeeCurrency
	responseObject.ID = symObject.BaseCurrency
	curerr, cur := GetRequest("https://api.hitbtc.com/api/2/public/currency/" + symObject.BaseCurrency)
	var curObject entity.CurrencyRequest
	json.Unmarshal(cur, &curObject)
	responseObject.FullName = curObject.FullName
	if curerr != nil || curObject.FullName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to Fetch Data From api.hitbtc Currency"))
	}
	responseObject.Symbol = ""
	h.Unlock()
	jsonBytes, err := json.Marshal(responseObject)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	logger.BootstrapLogger.Debug("Data Not Found")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Data not found"))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	logger.BootstrapLogger.Debug("Internal Server Error")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func GetRequest(uri string) (error, []byte) {
	logger.Logger.Debug("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Debug(err.Error())
		return err, bodyBytes
	}
	return nil, bodyBytes
}
