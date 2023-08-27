package storage

import (
	"context"
	"database/sql"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

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
