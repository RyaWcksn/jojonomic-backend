package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/transaction"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

type IService interface {
	Consume(ctx context.Context)
}

type ServiceImpl struct {
	brokerImpl      broker.IBroker
	storageImpl     storage.IStorage
	tarnsactionImpl transaction.ITransaction
	log             logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(b broker.IBroker, l logger.ILogger, s storage.IStorage, t transaction.ITransaction) *ServiceImpl {
	return &ServiceImpl{
		brokerImpl:      b,
		storageImpl:     s,
		tarnsactionImpl: t,
		log:             l,
	}
}
