package service

import (
	"github.com/bino1490/crypto-svc/pkg/entity"
	"github.com/bino1490/crypto-svc/pkg/logger"
	"github.com/bino1490/crypto-svc/pkg/repository"
)

type DBService interface {
	GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error)
}

//-- Service ----
type Service struct {
	repo repository.DbRepository
}

//--NewService  ----
func NewService(r repository.DbRepository) *Service {
	logger.BootstrapLogger.Debug("Entering Service.NewService() ...")
	return &Service{
		repo: r,
	}
}

//GetDBRecords redirects to repo layer to perform db operations
func (s *Service) GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error) {
	logger.Logger.Debug("Entering Service.GetDBRecords() ...")
	return s.repo.GetDBRecords(request)
}
