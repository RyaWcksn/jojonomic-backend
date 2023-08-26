package service

import (
	"context"
	"encoding/json"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/domain/storage"
)

// Consume implements IService.
func (s *ServiceImpl) Consume(ctx context.Context) {
	messages := s.brokerImpl.Consume(ctx)
	for message := range messages {
		s.log.Infof("Receive message := %v", string(message))
		var dbPayload storage.StorageEntity
		if err := json.Unmarshal(message, &dbPayload); err != nil {
			s.log.Errorf("Error while decode message := %v", err)
		}
		err := s.storageImpl.Insert(ctx, &dbPayload)
		if err != nil {
			s.log.Errorf("Error while decode message := %v", err)
		}
	}
}
