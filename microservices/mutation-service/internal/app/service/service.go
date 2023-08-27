package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
)

//go:generate mockgen -source service.go -destination service_mock.go -package service
type IService interface {
	FetchMutation(ctx context.Context, payload *dto.CheckMutasiReq) (res *[]storage.StorageEntity, err error)
}

type ServiceImpl struct {
	storageImpl storage.IStorage
	log         logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(s storage.IStorage, l logger.ILogger) *ServiceImpl {
	return &ServiceImpl{
		storageImpl: s,
		log:         l,
	}
}
