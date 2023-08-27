package service

import (
	"context"
	"errors"
	"math"

	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/constant"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/dto"
	rr "github.com/RyaWcksn/jojonomic-backend/topup-service/internal/errors"
)

// PublishMessage implements IService.
func (s *ServiceImpl) PublishMessage(ctx context.Context, payload *dto.TopupRequest) error {

	price, err := s.priceImpl.FetchPrice(ctx, payload.ReffId)
	if err != nil {
		s.log.Errorf("ERROR := %v, REFF_ID := %v", err, payload.ReffId)
		return err
	}

	if payload.Price != price.Data.HargaTopup {
		s.log.Errorf("ERROR := %s, REFF_ID := %v", "Harga kurang dari harga topup ", payload.ReffId)
		return rr.GetError(payload.ReffId, errors.New("harga tidak sama dengan harga topup"))
	}
	// Check if the amount is valid
	if payload.GoldWeight < 0.001 {
		s.log.Errorf("ERROR := %s, REFF_ID := %v", "Gram tidak valid", payload.ReffId)
		return rr.GetError(payload.ReffId, errors.New("Gram tidak valid"))
	}

	// Round the amount to the nearest multiple of 0.001
	roundedAmount := math.Round(payload.GoldWeight*1000) / 1000

	// Check if the rounded amount is the same as the original amount
	if roundedAmount != payload.GoldWeight {
		s.log.Errorf("ERROR := %s, REFF_ID := %v", "Gram tidak valid, harus berkelipatan 0.001", payload.ReffId)
		return rr.GetError(payload.ReffId, errors.New("Gram tidak valid, harus berkelipatan 0.001"))
	}

	brokerPayload := broker.BrokerMessage{
		ReffId:       payload.ReffId,
		Type:         constant.TOPUP,
		Norek:        payload.Norek,
		HargaTopup:   price.Data.HargaTopup,
		HargaBuyBack: price.Data.HargaBuyback,
		GoldWeight:   payload.GoldWeight,
	}

	err = s.brokerImpl.Publish(ctx, &brokerPayload)
	if err != nil {
		s.log.Errorf("Error while publish message", err)
		return err
	}

	return nil

}
