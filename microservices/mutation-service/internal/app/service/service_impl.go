package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/dto"
)

// FetchHarga implements IService.
func (s *ServiceImpl) FetchMutation(ctx context.Context, payload *dto.CheckMutasiReq) (res *[]storage.StorageEntity, err error) {

	mutationReq := storage.StorageRequest{
		From:   int64(payload.StartDate),
		To:     int64(payload.EndDate),
		ReffId: payload.ReffId,
		Norek:  payload.Norek,
	}
	harga, err := s.storageImpl.FetchMutation(ctx, &mutationReq)
	if err != nil {
		s.log.Errorf("Err := %v, REFF_ID := %v", err, payload.ReffId)
		return nil, err
	}

	return harga, nil
}
