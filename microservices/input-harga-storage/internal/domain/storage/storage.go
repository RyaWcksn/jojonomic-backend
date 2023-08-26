package storage

import (
	"context"
	"database/sql"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/logger"
)

type IStorage interface {
	Insert(ctx context.Context, payload *StorageEntity) error
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
