package storage

import (
	"context"
	"database/sql"

	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/pkgs/logger"
)

//go:generate mockgen -source storage.go -destination storage_mock.go -package storage
type IStorage interface {
	Get(ctx context.Context, payload *StorageEntityReq) (res *StorageEntityRes, err error)
}

type StorageImpl struct {
	sql *sql.DB
	log logger.ILogger
}

var _ IStorage = (*StorageImpl)(nil)

func NewStorage(sql *sql.DB, l logger.ILogger) *StorageImpl {
	return &StorageImpl{
		sql: sql,
		log: l,
	}
}
