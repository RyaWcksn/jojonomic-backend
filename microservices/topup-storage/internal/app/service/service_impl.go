package service

import (
	"context"
	"encoding/json"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/domain/broker"
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
		goldBalance.Data.GoldBalance += producedMessage.GoldWeight
		transactionPayload := transaction.TransactionEntity{
			ReffId:       producedMessage.ReffId,
			Type:         producedMessage.Type,
			Norek:        producedMessage.Norek,
			HargaTopup:   producedMessage.HargaTopup,
			HargaBuyBack: producedMessage.HargaBuyBack,
			GoldWeight:   producedMessage.GoldWeight,
			// Still not sure is the transaction will be sum of current balance and topup
			// Or last balance, let's assume sum for now
			//GoldBalance:  goldBalance.GoldBalance,
			GoldBalance: goldBalance.Data.GoldBalance,
		}
		err = s.tarnsactionImpl.Insert(ctx, &transactionPayload)
		if err != nil {
			s.log.Errorf("Error while insert transaction := %v", err)
			return
		}
	}
}
