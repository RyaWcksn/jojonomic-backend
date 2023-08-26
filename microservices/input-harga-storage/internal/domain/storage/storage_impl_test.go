package storage

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-storage/internal/logger"
)

func TestStorageImpl_Insert(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	l := logger.New("", "", "")

	type args struct {
		ctx     context.Context
		payload *StorageEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: &StorageEntity{
					AdminId:      "a0001",
					ReffId:       "refId123",
					HargaTopup:   910000,
					HargaBuyback: 820000,
				},
			},
			wantErr: false,
		},
		{
			name: "Duplicate",
			args: args{
				ctx: context.Background(),
				payload: &StorageEntity{
					AdminId:      "a0001",
					ReffId:       "refId123",
					HargaTopup:   910000,
					HargaBuyback: 820000,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			query := "INSERT INTO tbl_harga\\(reff_id, admin_id, harga_topup, harga_buyback\\) VALUES\\(\\$1, \\$2, \\$3, \\$4\\)"
			if tt.wantErr {
				mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(tt.args.payload.ReffId, tt.args.payload.AdminId, tt.args.payload.HargaTopup, tt.args.payload.HargaBuyback).
					WillReturnError(fmt.Errorf("duplicate entry"))
			} else {
				mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(tt.args.payload.ReffId, tt.args.payload.AdminId, tt.args.payload.HargaTopup, tt.args.payload.HargaBuyback).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}
			mock.ExpectCommit()

			s := NewStorage(db, l)
			if err := s.Insert(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("StorageImpl.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
