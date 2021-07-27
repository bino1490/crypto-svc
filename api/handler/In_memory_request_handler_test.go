package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bino1490/crypto-svc/pkg/service"
	gomock "github.com/golang/mock/gomock"
)

func TestCurreny_Success(t *testing.T) {

	// Initialize
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSer := service.NewMockMemoryService(mockCtrl)

	req := httptest.NewRequest(http.MethodGet, "/currency/all", nil)
	res := httptest.NewRecorder()
	mockSer.EXPECT().InMemGetPOST(res, req).AnyTimes().Do(req)

	handler := http.Handler(InCurrencyReqHandler(service.NewInMemService()))
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
