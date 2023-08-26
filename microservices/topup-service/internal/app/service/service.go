package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/domain/price"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/logger"
)

//go:generate mockgen -source service.go -destination service_mock.go -package service
type IService interface {
	PublishMessage(ctx context.Context, payload *dto.TopupRequest) error
}

type ServiceImpl struct {
	priceImpl  price.IPrice
	brokerImpl broker.IBroker
	log        logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(broker broker.IBroker, l logger.ILogger, p price.IPrice) *ServiceImpl {
	return &ServiceImpl{
		priceImpl:  p,
		brokerImpl: broker,
		log:        l,
	}
}
