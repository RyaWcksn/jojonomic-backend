package storage

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

type IStorage interface {
	Get(ctx context.Context, payload *StorageEntityReq) (res *SaldoEntity, err error)
}

type StorageImpl struct {
	cfg config.Config
	log logger.ILogger
}

var _ IStorage = (*StorageImpl)(nil)

func NewStorage(cfg config.Config, l logger.ILogger) *StorageImpl {
	return &StorageImpl{
		cfg: cfg,
		log: l,
	}
}
