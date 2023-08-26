package storage

import (
	"context"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/errors"
	"github.com/teris-io/shortid"
)

// FetchHarga implements IStorage.
func (s *StorageImpl) FetchHarga(ctx context.Context) (res *StorageEntity, err error) {
	sqlCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT harga_topup, harga_buyback FROM tbl_harga ORDER BY created_at desc"

	row := s.sql.QueryRowContext(sqlCtx, query)
	var payload StorageEntity
	err = row.Scan(&payload.HargaTopup, &payload.HargaBuyback)
	if err != nil {
		refId, _ := shortid.Generate()
		s.l.Errorf("Error while scan := %v", err)
		return nil, errors.GetError(refId, err)
	}

	return &payload, nil

}
