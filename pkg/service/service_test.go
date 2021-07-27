package service

import (
	reflect "reflect"
	"testing"
	"time"

	"github.com/bino1490/crypto-svc/pkg/entity"
	"github.com/bino1490/crypto-svc/pkg/repository"
	gomock "github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestNewService(t *testing.T) {
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
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDBRecords(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDb := repository.NewMockDbRepository(mockCtrl)
	testService := NewService(mockDb)

	sampleEntity := entity.DBRequest{
		EndDate:   "2021-05-01",
		StartDate: "2021-10-02",
		MaxCount:  1,
		MinCount:  2,
	}

	sampleResp := entity.DBRecord{
		Key:        "1",
		CreatedAt:  time.Now(),
		TotalCount: 15,
	}
	var rsps []entity.DBRecord
	rsps = append(rsps, sampleResp)
	//Mocking
	mockDb.EXPECT().GetDBRecords(sampleEntity).Return(rsps, nil)
	//Call invocation
	result, _ := testService.GetDBRecords(sampleEntity)
	//Assertion
	assert.Equal(t, rsps, result)
}
