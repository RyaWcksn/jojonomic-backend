package transaction

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/jojonomic-backend/topup-storage/internal/logger"
)

func TestTransactionImpl_Insert(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	l := logger.New("", "", "")
	type args struct {
		ctx     context.Context
		payload *TransactionEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				payload: &TransactionEntity{
					ReffId:       "ref123",
					Type:         "topup",
					Norek:        "rek123",
					HargaTopup:   910000,
					HargaBuyBack: 820000,
					GoldWeight:   0.01,
					GoldBalance:  0.01,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock.ExpectBegin()
			query := "INSERT INTO tbl_transaksi\\(reff_id, norek, type, harga_topup, harga_buyback, gold_weight, gold_balance, created_at\\) VALUES\\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8\\)"
			if tt.wantErr {
				mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(tt.args.payload.ReffId, tt.args.payload.Norek, tt.args.payload.Type, tt.args.payload.HargaTopup, tt.args.payload.HargaBuyBack, tt.args.payload.GoldWeight, tt.args.payload.GoldBalance, time.Now().Local().Unix).
					WillReturnError(fmt.Errorf("duplicate entry"))
			} else {
				mock.ExpectPrepare(query).
					ExpectExec().
					WithArgs(tt.args.payload.ReffId, tt.args.payload.Norek, tt.args.payload.Type, tt.args.payload.HargaTopup, tt.args.payload.HargaBuyBack, tt.args.payload.GoldWeight, tt.args.payload.GoldBalance, time.Now().Local().Unix).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}
			mock.ExpectCommit()
			s := NewTransaction(db, l)
			if err := s.Insert(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("TransactionImpl.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
