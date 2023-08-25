package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/logger"
)

//go:generate mockgen -source service.go -destination service_mock.go -package service
type IService interface {
	PublishMessage(ctx context.Context, payload *dto.InputHargaRequest) error
}

type ServiceImpl struct {
	brokerImpl broker.IBroker
	log        logger.ILogger
}

var _ IService = (*ServiceImpl)(nil)

func NewService(broker broker.IBroker, l logger.ILogger) *ServiceImpl {
	return &ServiceImpl{
		brokerImpl: broker,
		log:        l,
	}
}
