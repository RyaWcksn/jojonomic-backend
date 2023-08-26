package storage

import (
	"context"
	"time"
)

// Insert implements IStorage.
func (s *StorageImpl) Insert(ctx context.Context, payload *StorageEntity) error {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	tx, err := s.sql.Begin()
	if err != nil {
		s.log.Errorf("[ERR] While starting transaction := %v", err)
	}
	query := "INSERT INTO tbl_harga(reff_id, admin_id, harga_topup, harga_buyback) VALUES($1, $2, $3, $4)"
	stmt, err := tx.PrepareContext(ctxDb, query)
	if err != nil {
		s.log.Errorf("[ERR] While prepare statement := %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.ReffId, payload.AdminId, payload.HargaTopup, payload.HargaBuyback)
	if err != nil {
		tx.Rollback()
		s.log.Errorf("[ERR] While executing query := %v", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		s.log.Errorf("[ERR] While commit transaction := %v", err)

	}
	return nil

}
