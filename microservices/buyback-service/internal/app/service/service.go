package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/price"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/logger"
)

type IService interface {
	PublishMessage(ctx context.Context, payload *dto.BuybackRequest) error
}

type ServiceImpl struct {
	saldoImpl  storage.IStorage
	brokerImpl broker.IBroker
	priceImpl  price.IPrice
	log        logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(broker broker.IBroker, l logger.ILogger, s storage.IStorage, p price.IPrice) *ServiceImpl {
	return &ServiceImpl{
		saldoImpl:  s,
		brokerImpl: broker,
		priceImpl:  p,
		log:        l,
	}
}
