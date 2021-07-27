package repository

import "github.com/bino1490/crypto-svc/pkg/entity"

type DbRepository interface {
	GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error)
}
