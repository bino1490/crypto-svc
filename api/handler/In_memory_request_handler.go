package handler

import (
	"net/http"

	"github.com/bino1490/crypto-svc/pkg/logger"
	"github.com/bino1490/crypto-svc/pkg/service"
)

// InMemReqHandler the Handler Layer for future business logic enhancement
// for now just redirecting to Service Layer
func InMemReqHandler(inMemSvc *service.MemHandlers) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Logger.Debug("Entering handler.InMemReqHandler() ...")
		// redirects to service layer
		inMemSvc.InMemGetPOST(w, r)
		return
	})

}
