package storage

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
)

func TestStorageImpl_FetchHarga(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	l := logger.New("", "", "")

	rowLists := []string{"created_at", "type", "gold_weight", "harga_topup", "harga_buyback", "gold_balance"}
	query := "SELECT created_at, type, gold_weight, harga_topup, harga_buyback, gold_balance FROM tbl_transaksi WHERE norek = \\$1 AND created_at BETWEEN \\$2 AND \\$3 ORDER BY created_at DESC;"

	type args struct {
		ctx     context.Context
		payload *StorageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantRes  *[]StorageEntity
		wantErr  bool
		wantMock func()
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				payload: &StorageRequest{
					From:   1693092631,
					To:     1693092658,
					ReffId: "ref123",
					Norek:  "rek001",
				},
			},
			wantRes: &[]StorageEntity{
				{
					Date:         1693092631,
					Type:         "topup",
					GoldWeight:   0.001,
					HargaTopup:   910000,
					HargaBuyback: 820000,
					GoldBalance:  0.002,
				},
				{
					Date:         1693092658,
					Type:         "topup",
					GoldWeight:   0.001,
					HargaTopup:   1000,
					HargaBuyback: 500,
					GoldBalance:  0.003,
				},
			},
			wantErr: false,
			wantMock: func() {
				row := mock.NewRows(rowLists).
					AddRow(1693092631, "topup", 0.001, 910000, 820000, 0.002).
					AddRow(1693092658, "topup", 0.001, 1000, 500, 0.003)
				mock.ExpectQuery(query).
					WithArgs("rek001", 1693092631, 1693092658).
					WillReturnRows(row)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewStorage(db, l)
			gotRes, err := s.FetchMutation(tt.args.ctx, tt.args.payload)
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
