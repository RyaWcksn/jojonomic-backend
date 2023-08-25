package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/dto"
)

// PublishMessage implements IService.
func (s *ServiceImpl) PublishMessage(ctx context.Context, payload *dto.InputHargaRequest) error {

	brokerPayload := broker.BrokerMessage{
		AdminId:      payload.AdminId,
		ReffId:       payload.ReffID,
		HargaTopup:   payload.HargaTopup,
		HargaBuyback: payload.HargaBuyback,
	}

	err := s.brokerImpl.Publish(ctx, &brokerPayload)
	if err != nil {
		s.log.Errorf("Error while publish message", err)
		return err
	}

	return nil
}
