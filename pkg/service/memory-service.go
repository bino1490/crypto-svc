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
	getRe      = regexp.MustCompile(`^\/in-memory[\/]*$`)
	getByKeyRe = regexp.MustCompile(`^\/in-memory\/[^*]`)
	createRe   = regexp.MustCompile(`^\/in-memory[\/]*$`)
)

//creating interface with abstract methods
type MemoryService interface {
	InMemGetPOST(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetMemData(w http.ResponseWriter, r *http.Request)
	GetDataByKey(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type MemHandlers struct {
	sync.Mutex
	store map[string]entity.InMemoryRequest
}

func NewInMemService() *MemHandlers {
	return &MemHandlers{
		store: map[string]entity.InMemoryRequest{"active-tabs": entity.InMemoryRequest{
			Key: "active-tabs", Value: "getir"}},
	}

}

//InMemGetPOST to perform in memory Add/GET operations
func (h *MemHandlers) InMemGetPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && getRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createRe.MatchString(r.URL.Path):
		h.Post(w, r)
		return
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
	reqs := make([]entity.InMemoryRequest, len(h.store))

	h.Lock()
	i := 0
	for _, data := range h.store {
		reqs[i] = data
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(reqs)
	if err != nil {
		logger.BootstrapLogger.Error("Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
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
	u, ok := h.store[matches[2]]
	h.Unlock()
	if !ok {
		logger.BootstrapLogger.Error("Value not found in memory")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Value not found in Memory"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//Add the data to in memory
func (h *MemHandlers) Post(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debug("Entering service.POST() ...")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.BootstrapLogger.Error("Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		logger.BootstrapLogger.Error("need content-type 'application/json")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var memReq entity.InMemoryRequest
	err = json.Unmarshal(bodyBytes, &memReq)
	if err != nil || memReq.Key == "" || memReq.Value == "" {
		logger.BootstrapLogger.Error("Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Input Request"))
		return
	}
	h.Lock()
	h.store[memReq.Key] = memReq
	defer h.Unlock()
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	logger.BootstrapLogger.Error("Data Not Found")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Data not found"))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	logger.BootstrapLogger.Error("Internal Server Error")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}
