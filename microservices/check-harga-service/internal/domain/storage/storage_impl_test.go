package storage

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/logger"
)

func TestStorageImpl_FetchHarga(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	l := logger.New("", "", "")

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantRes *StorageEntity
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
			},
			wantRes: &StorageEntity{
				HargaTopup:   910000,
				HargaBuyback: 820000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rowLists := []string{"harga_topup", "harga_buyback"}
			query := "SELECT harga_topup, harga_buyback FROM tbl_harga ORDER BY created_at desc"
			if tt.wantRes != nil {
				// Set up the expected query and result
				rows := sqlmock.NewRows(rowLists).
					AddRow(tt.wantRes.HargaTopup, tt.wantRes.HargaBuyback)
				mock.ExpectQuery(query).
					WillReturnRows(rows)
			} else {
				// Set up the expected query with an empty result set
				mock.ExpectQuery(query).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{}))
			}
			s := NewStorage(db, l)
			gotRes, err := s.FetchHarga(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageImpl.FetchHarga() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("StorageImpl.FetchHarga() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
