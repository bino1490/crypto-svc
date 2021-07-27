package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bino1490/crypto-svc/pkg/entity"
	"github.com/bino1490/crypto-svc/pkg/service"
	gomock "github.com/golang/mock/gomock"
)

func TestDBRequest_Success(t *testing.T) {

	// Initialize
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSer := service.NewMockDBService(mockCtrl)
	dbrec := entity.DBRecord{
		Key:        "1",
		CreatedAt:  time.Now(),
		TotalCount: 10,
	}
	drreq := entity.DBRequest{
		EndDate:   "2025-01-01",
		StartDate: "2020-01-01",
		MaxCount:  100,
		MinCount:  1,
	}
	str, _ := json.Marshal(drreq)
	resourceReader := strings.NewReader(string(str))
	req := httptest.NewRequest(http.MethodPost, "/records", resourceReader)
	req.Header.Set("content-type", "application/json")
	res := httptest.NewRecorder()

	var recs []entity.DBRecord
	recs = append(recs, dbrec)
	mockSer.EXPECT().GetDBRecords(drreq).Return(recs, nil)

	handler := http.Handler(DBHandler(mockSer))
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
