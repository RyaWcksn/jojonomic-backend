package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/logger"
)

type IService interface {
	Consume(ctx context.Context)
}

type ServiceImpl struct {
	brokerImpl  broker.IBroker
	storageImpl storage.IStorage
	log         logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(b broker.IBroker, l logger.ILogger, s storage.IStorage) *ServiceImpl {
	return &ServiceImpl{
		brokerImpl:  b,
		storageImpl: s,
		log:         l,
	}
}
