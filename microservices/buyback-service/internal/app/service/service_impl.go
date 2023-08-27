package service

import (
	"context"
	"errors"
	"math"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/constant"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/dto"
	rr "github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/errors"
)

// PublishMessage implements IService.
func (s *ServiceImpl) PublishMessage(ctx context.Context, payload *dto.BuybackRequest) error {
	saldoPayload := storage.StorageEntityReq{
		Norek:  payload.Norek,
		ReffId: payload.ReffId,
	}
	saldo, err := s.saldoImpl.Get(ctx, &saldoPayload)
	if err != nil {
		return err
	}
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

	if saldo.Data.GoldBalance < payload.GoldWeight {
		s.log.Errorf("ERR := %s REFF_ID := %s ", "Saldo emas tidak cukup", payload.ReffId)
		return rr.GetError(payload.ReffId, errors.New("Saldo emas tidak mencukupi"))
	}

	price, err := s.priceImpl.FetchPrice(ctx, payload.ReffId)
	if err != nil {
		return err
	}

	brokerPayload := broker.BrokerMessage{
		ReffId:       payload.ReffId,
		Type:         constant.BUYBACK,
		Norek:        payload.Norek,
		HargaTopup:   price.Data.HargaTopup,
		HargaBuyBack: price.Data.HargaBuyback,
		GoldWeight:   payload.GoldWeight,
	}

	if price.Data.HargaBuyback != payload.Price {
		s.log.Errorf("ERR := %s REFF_ID := %s ", "harga tidak sama", payload.ReffId)
		return rr.GetError(payload.ReffId, errors.New("harga tidak sama"))
	}

	err = s.brokerImpl.Publish(ctx, &brokerPayload)
	if err != nil {
		s.log.Errorf("Error while publish message", err)
		return err
	}

	return nil
}
