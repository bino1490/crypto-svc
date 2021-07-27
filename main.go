//Author: Bino Patric Prakah G
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bino1490/crypto-svc/api/handler"
	"github.com/bino1490/crypto-svc/pkg/config"
	"github.com/bino1490/crypto-svc/pkg/logger"
	"github.com/bino1490/crypto-svc/pkg/service"
)

//-- Main to Inatialize the service
func main() {
	logger.BootstrapLogger.Info("Service starting...")
	initService()
}

//-- initService initialize the service ----
func initService() {
	logger.BootstrapLogger.Debug("Entering initService...")
	logger.BootstrapLogger.Info("Starting " + config.SrvConfig.GetString("application.name") +
		" with profile=" + config.SrvConfig.GetString("profile") + " properties")

	//Handler for InMemory Services
	inMemSvc := service.NewInMemService()
	addInMemHandlers(inMemSvc)
	//err := http.ListenAndServe(":"+config.SrvConfig.GetString("http.port"), nil)
	//os.Setenv("PORT", "5000")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		logger.BootstrapLogger.Error("Failed to ListenPort 5000")
		panic(err)
	}
}

//-- This is a primitive test handler that should be removed by the developer
func simplehandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		title := "simple response..."
		fmt.Fprintf(w, "Hello from:  "+title+"\n")
	})
}

//addInMemHandlers to hadle the business operations
func addInMemHandlers(inMemSvc *service.MemHandlers) {
	http.Handle("/currency/", handler.InCurrencyReqHandler(inMemSvc))
	http.Handle("/", simplehandler())
}
