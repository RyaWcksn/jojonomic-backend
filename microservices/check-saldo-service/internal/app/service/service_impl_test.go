package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/pkgs/logger"
	gomock "github.com/golang/mock/gomock"
)

func TestServiceImpl_FetchSaldo(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := storage.NewMockIStorage(ctrl)
	l := logger.New("", "", "")
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		payload *dto.CheckSaldoReq
	}
	tests := []struct {
		name     string
		args     args
		wantRes  *storage.StorageEntityRes
		wantErr  bool
		wantMock func()
	}{

		{
			name: "Success",
			args: args{
				ctx: ctx,
				payload: &dto.CheckSaldoReq{
					Norek:  "rek123",
					ReffId: "ref123",
				},
			},
			wantRes: &storage.StorageEntityRes{Norek: "rek123", GoldBalance: 0.001},
			wantErr: false,
			wantMock: func() {
				storageMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&storage.StorageEntityRes{Norek: "rek123", GoldBalance: 0.001}, nil)
			},
		},
		{
			name: "Failed",
			args: args{
				ctx:     ctx,
				payload: &dto.CheckSaldoReq{},
			},
			wantRes: nil,
			wantErr: true,
			wantMock: func() {
				storageMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("err"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewService(storageMock, l)
			gotRes, err := s.FetchSaldo(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.FetchSaldo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ServiceImpl.FetchSaldo() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
