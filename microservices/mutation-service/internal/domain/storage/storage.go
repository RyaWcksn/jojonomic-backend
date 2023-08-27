package storage

import (
	"context"
	"database/sql"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
)

//go:generate mockgen -source storage.go -destination storage_mock.go -package storage
type IStorage interface {
	FetchMutation(ctx context.Context, payload *StorageRequest) (res *[]StorageEntity, err error)
}

type StorageImpl struct {
	sql *sql.DB
	l   logger.ILogger
}

var _ IStorage = (*StorageImpl)(nil)

func NewStorage(sql *sql.DB, l logger.ILogger) *StorageImpl {
	return &StorageImpl{
		sql: sql,
		l:   l,
	}
}
