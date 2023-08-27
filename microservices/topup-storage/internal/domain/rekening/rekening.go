package rekening

import (
	"context"
	"database/sql"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

type IRekening interface {
	UpdateSaldo(ctx context.Context, payload *RekeningEntity) error
}

type RekeningImpl struct {
	sql *sql.DB
	l   logger.ILogger
}

// UpdateSaldo implements IRekening.
func (r *RekeningImpl) UpdateSaldo(ctx context.Context, payload *RekeningEntity) error {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := r.sql.Begin()
	if err != nil {
		r.l.Errorf("[ERR] While starting transaction := %v", err)
	}
	query := "UPDATE tbl_rekening SET gold_balance = $1 WHERE norek = $2"
	stmt, err := tx.PrepareContext(ctxDb, query)
	if err != nil {
		r.l.Errorf("[ERR] While prepare statement := %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.GoldWeight, payload.Norek)
	if err != nil {
		tx.Rollback()
		r.l.Errorf("[ERR] While executing query := %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		r.l.Errorf("[ERR] While commit transaction := %v", err)

	}
	return nil

}

var _ IRekening = (*RekeningImpl)(nil)

func NewRekening(sql *sql.DB, l logger.ILogger) *RekeningImpl {
	return &RekeningImpl{
		sql: sql,
		l:   l,
	}
}
