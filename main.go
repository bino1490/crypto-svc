//Author: Bino Patric Prakah G
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bino1490/crypto-svc/api/handler"
	"github.com/bino1490/crypto-svc/pkg/config"
	"github.com/bino1490/crypto-svc/pkg/entity"
	"github.com/bino1490/crypto-svc/pkg/logger"
	"github.com/bino1490/crypto-svc/pkg/repository"
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

	// handler for DB Services
	repository := initDatabase()
	dbservice := service.NewService(repository)
	addDBMemHandlers(dbservice)

	//Handler for InMemory Services
	inMemSvc := service.NewInMemService()
	addInMemHandlers(inMemSvc)
	//err := http.ListenAndServe(":"+config.SrvConfig.GetString("http.port"), nil)
	//os.Setenv("PORT", "5000")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		logger.BootstrapLogger.Error("Failed to ListenPort 8080")
		panic(err)
	}
}

//-- SYSTEM GENEREATED: Testing Code ----
//-- This is a primitive test handler that should be removed by the developer
func simplehandler(service service.DBService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		title := "simple response..."
		fmt.Fprintf(w, "Hello from:  "+title+"\n")
	})
}

//to hadle the business operations
func addDBMemHandlers(dbSvc service.DBService) {
	//-- This is a primitive test handler that should be removed by the developer
	http.Handle("/", simplehandler(dbSvc))
	http.Handle("/records", handler.DBHandler(dbSvc))
}

//addInMemHandlers to hadle the business operations
func addInMemHandlers(inMemSvc *service.MemHandlers) {
	//http.HandleFunc("/in-memory", inMemSvc.InMemGetPOST)
	http.Handle("/in-memory", handler.InMemReqHandler(inMemSvc))
	http.Handle("/in-memory/", handler.InMemReqHandler(inMemSvc))
}

//-- To initialize the database ----
func initDatabase() repository.DbRepository {
	logger.BootstrapLogger.Debug("Entering initDatabase...")

	if config.SrvConfig.GetString(
		"database.mongodb.enable") == "true" {
		logger.BootstrapLogger.Debug("About to initialize MongoDB repo...")
		return repository.NewCbRepository()

	} else {
		// Throw panic
		logger.BootstrapLogger.Error("Incorrect database configuration settings. Can't proceed!")
		panic(entity.ErrInvalidConfig)
	}
}
