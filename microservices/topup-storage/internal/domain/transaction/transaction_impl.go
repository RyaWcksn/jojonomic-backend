package transaction

import (
	"context"
	"time"
)

// Insert implements ITransaction.
func (s *TransactionImpl) Insert(ctx context.Context, payload *TransactionEntity) error {

	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := s.sql.Begin()
	if err != nil {
		s.log.Errorf("[ERR] While starting transaction := %v", err)
	}
	currentTime := time.Now().Local().Unix()
	query := "INSERT INTO tbl_transaksi(reff_id, norek, type, harga_topup, harga_buyback, gold_weight, gold_balance, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	stmt, err := tx.PrepareContext(ctxDb, query)
	if err != nil {
		s.log.Errorf("[ERR] While prepare statement := %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.ReffId, payload.Norek, payload.Type, payload.HargaTopup, payload.HargaBuyBack, payload.GoldWeight, payload.GoldBalance, currentTime)
	if err != nil {
		tx.Rollback()
		s.log.Errorf("[ERR] While executing query := %v", err)
	}

	err = tx.Commit()
	if err != nil {
		s.log.Errorf("[ERR] While commit transaction := %v", err)

	}
	return nil
}
