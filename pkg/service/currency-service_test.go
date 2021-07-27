package service

import (
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewMemService(t *testing.T) {
	tests := []struct {
		name string
		//args args
		want *MemoryService
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

	req := httptest.NewRequest(http.MethodGet, "/currency/test", nil)
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

	req := httptest.NewRequest(http.MethodGet, "/currency/all", nil)
	res := httptest.NewRecorder()
	testService.InMemGetPOST(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}
