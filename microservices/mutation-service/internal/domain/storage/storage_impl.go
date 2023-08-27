package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/errors"
)

// FetchHarga implements IStorage.
func (s *StorageImpl) FetchMutation(ctx context.Context, payload *StorageRequest) (res *[]StorageEntity, err error) {
	sqlCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT created_at, type, gold_weight, harga_topup, harga_buyback, gold_balance FROM tbl_transaksi WHERE norek = $1 AND created_at BETWEEN $2 AND $3 ORDER BY created_at DESC;"

	s.l.Infof("From date %v", payload.From)
	s.l.Infof("To date %v", payload.To)
	s.l.Infof("Norek %v", payload.Norek)

	rows, err := s.sql.QueryContext(sqlCtx, query, payload.Norek, payload.From, payload.To)
	if err != nil {
		s.l.Errorf("Error while scan := %v REFF_id := %v", err, payload.ReffId)
		return nil, errors.GetError(payload.ReffId, err)
	}
	defer rows.Close()

	var Mutations []StorageEntity
	for rows.Next() {
		var mutation StorageEntity
		err := rows.Scan(
			&mutation.Date,
			&mutation.Type,
			&mutation.GoldWeight,
			&mutation.HargaTopup,
			&mutation.HargaBuyback,
			&mutation.GoldBalance,
		)
		if err != nil {
			s.l.Errorf("Error while scanning row: %v", err)
			return nil, errors.GetError(payload.ReffId, err)
		}

		Mutations = append(Mutations, mutation)
	}

	fmt.Println(Mutations)
	return &Mutations, nil

}
