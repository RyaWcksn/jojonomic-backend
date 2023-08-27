package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/domain/rekening"
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/domain/transaction"
	"github.com/RyaWcksn/jojonomic-backend/buyback-storage/internal/pkgs/logger"
)

type IService interface {
	Consume(ctx context.Context)
}

type ServiceImpl struct {
	rekeningImpl    rekening.IRekening
	brokerImpl      broker.IBroker
	storageImpl     storage.IStorage
	tarnsactionImpl transaction.ITransaction
	log             logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(b broker.IBroker, l logger.ILogger, s storage.IStorage, t transaction.ITransaction, r rekening.IRekening) *ServiceImpl {
	return &ServiceImpl{
		rekeningImpl:    r,
		brokerImpl:      b,
		storageImpl:     s,
		tarnsactionImpl: t,
		log:             l,
	}
}
