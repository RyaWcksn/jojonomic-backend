package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/domain/storage"
)

// FetchHarga implements IService.
func (s *ServiceImpl) FetchHarga(ctx context.Context) (res *storage.StorageEntity, err error) {
	harga, err := s.storageImpl.FetchHarga(ctx)
	if err != nil {
		s.log.Errorf("Err := %v", err)
		return nil, err
	}

	return harga, nil
}
