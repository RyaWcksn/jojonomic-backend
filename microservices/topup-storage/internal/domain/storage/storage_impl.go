package storage

import (
	"context"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/errors"
)

// Get implements IStorage.
func (s *StorageImpl) Get(ctx context.Context, payload *StorageEntityReq) (res *StorageEntityRes, err error) {
	ctxDb, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var recDetail StorageEntityRes

	query := "SELECT gold_balance FROM tbl_rekening WHERE norek = $1"
	err = s.sql.QueryRowContext(ctxDb, query, payload.Norek).Scan(&recDetail.GoldBalance)
	if err != nil {
		s.log.Errorf("ERR := %v, Reff_id := %v", err, payload.ReffId)
		return nil, errors.GetError(payload.ReffId, err)
	}

	return &recDetail, nil
}
