package transaction

import (
	"context"
	"database/sql"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

type ITransaction interface {
	Insert(ctx context.Context, payload *TransactionEntity) error
}

type TransactionImpl struct {
	sql *sql.DB
	log logger.ILogger
}

var _ ITransaction = (*TransactionImpl)(nil)

func NewTransaction(sql *sql.DB, l logger.ILogger) *TransactionImpl {
	return &TransactionImpl{
		sql: sql,
		log: l,
	}
}
