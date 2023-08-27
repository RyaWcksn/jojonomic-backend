package storage

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/pkgs/logger"
)

func TestStorageImpl_Get(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	l := logger.New("", "", "")

	type args struct {
		ctx     context.Context
		payload *StorageEntityReq
	}
	tests := []struct {
		name     string
		args     args
		wantRes  *StorageEntityRes
		wantErr  bool
		wantMock func()
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				payload: &StorageEntityReq{
					Norek:  "rek123",
					ReffId: "ref123",
				},
			},
			wantRes: &StorageEntityRes{
				GoldBalance: 0.01,
			},
			wantErr: false,
			wantMock: func() {
				rows := sqlmock.NewRows([]string{"gold_balance"}).
					AddRow(0.01)
				mock.ExpectQuery("SELECT gold_balance FROM tbl_rekening WHERE norek = \\$1").
					WithArgs("rek123").
					WillReturnRows(rows)
			},
		},
		{
			name: "Error on query",
			args: args{
				ctx: context.TODO(),
				payload: &StorageEntityReq{
					Norek: "54321",
				},
			},
			wantMock: func() {
				mock.ExpectQuery("SELECT gold_balance FROM tbl_rekening WHERE norek = ?").
					WithArgs("54321").
					WillReturnError(sql.ErrNoRows) // Replace with your desired error
			},
			wantRes: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewStorage(db, l)
			gotRes, err := s.Get(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("StorageImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("StorageImpl.Get() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
