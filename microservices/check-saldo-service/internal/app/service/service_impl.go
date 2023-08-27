package service

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/dto"
)

func (s *ServiceImpl) FetchSaldo(ctx context.Context, payload *dto.CheckSaldoReq) (res *storage.StorageEntityRes, err error) {

	saldoPayload := storage.StorageEntityReq{
		Norek:  payload.Norek,
		ReffId: payload.ReffId,
	}
	saldo, err := s.storageImpl.Get(ctx, &saldoPayload)
	if err != nil {
		s.log.Errorf("Err := %v", err)
		return nil, err
	}

	return saldo, nil
}
