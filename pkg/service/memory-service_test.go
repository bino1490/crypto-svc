package service

import (
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	"github.com/bino1490/crypto-svc/pkg/repository"
	gomock "github.com/golang/mock/gomock"
)

func TestNewMemService(t *testing.T) {
	type args struct {
		r repository.DbRepository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRecords(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testService := NewInMemService()

	req := httptest.NewRequest(http.MethodGet, "/in-memory", nil)
	res := httptest.NewRecorder()
	testService.Get(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}

func TestInMemGetPOST(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testService := NewInMemService()

	req := httptest.NewRequest(http.MethodGet, "/in-memory", nil)
	res := httptest.NewRecorder()
	testService.InMemGetPOST(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}

func TestPOSTNegative(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testService := NewInMemService()

	req := httptest.NewRequest(http.MethodPost, "/in-memory", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	testService.Post(res, req)
	if res.Code != http.StatusBadRequest {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}
