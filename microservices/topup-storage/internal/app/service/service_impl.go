package service

import (
	"context"
	"encoding/json"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/rekening"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/transaction"
)

// Consume implements IService.
func (s *ServiceImpl) Consume(ctx context.Context) {
	messages := s.brokerImpl.Consume(ctx)
	for message := range messages {
		s.log.Infof("Receive message := %v", string(message))
		var producedMessage broker.BrokerEntity
		if err := json.Unmarshal(message, &producedMessage); err != nil {
			s.log.Errorf("Error while decode message := %v", err)
		}
		goldBalancePayload := storage.StorageEntityReq{
			Norek:  producedMessage.Norek,
			ReffId: producedMessage.ReffId,
		}
		goldBalance, err := s.storageImpl.Get(ctx, &goldBalancePayload)
		if err != nil {
			s.log.Errorf("Error while get gold balance := %v", err)
			continue
		}
		currGoldBalance := goldBalance.Data.GoldBalance
		currGoldBalance += producedMessage.GoldWeight
		s.log.Infof("GOLD BALANCE := ", goldBalance.Data.GoldBalance)
		transactionPayload := transaction.TransactionEntity{
			ReffId:       producedMessage.ReffId,
			Type:         producedMessage.Type,
			Norek:        producedMessage.Norek,
			HargaTopup:   producedMessage.HargaTopup,
			HargaBuyBack: producedMessage.HargaBuyBack,
			GoldWeight:   producedMessage.GoldWeight,
			GoldBalance:  currGoldBalance,
		}
		err = s.tarnsactionImpl.Insert(ctx, &transactionPayload)
		if err != nil {
			s.log.Errorf("Error while insert transaction := %v", err)
			return
		}
		rekeningPayload := rekening.RekeningEntity{
			Norek:      producedMessage.Norek,
			GoldWeight: currGoldBalance,
		}
		err = s.rekeningImpl.UpdateSaldo(ctx, &rekeningPayload)
		if err != nil {
			s.log.Errorf("Error while insert transaction := %v", err)
			return
		}
	}
}
