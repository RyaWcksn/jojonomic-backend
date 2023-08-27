package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestServiceImpl_FetchHarga(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := storage.NewMockIStorage(ctrl)
	l := logger.New("", "", "")
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		payload *dto.CheckMutasiReq
	}

	tests := []struct {
		name     string
		args     args
		wantRes  *[]storage.StorageEntity
		wantErr  bool
		wantMock func()
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				payload: &dto.CheckMutasiReq{
					Norek:     "req123",
					StartDate: 123123213231,
					EndDate:   123213213123,
				},
			},
			wantRes: &[]storage.StorageEntity{{Date: 1693092631, Type: "topup", GoldWeight: 0.001, HargaTopup: 910000, HargaBuyback: 820000, GoldBalance: 0.002}, {Date: 1693092658, Type: "topup", GoldWeight: 0.001, HargaTopup: 1000, HargaBuyback: 500, GoldBalance: 0.003}},
			wantErr: false,
			wantMock: func() {
				storageMock.EXPECT().FetchMutation(gomock.Any(), gomock.Any()).Return(&[]storage.StorageEntity{{Date: 1693092631, Type: "topup", GoldWeight: 0.001, HargaTopup: 910000, HargaBuyback: 820000, GoldBalance: 0.002}, {Date: 1693092658, Type: "topup", GoldWeight: 0.001, HargaTopup: 1000, HargaBuyback: 500, GoldBalance: 0.003}}, nil)
			},
		},
		{
			name: "Failed",
			args: args{
				ctx:     ctx,
				payload: &dto.CheckMutasiReq{},
			},
			wantRes: nil,
			wantErr: true,
			wantMock: func() {
				storageMock.EXPECT().FetchMutation(gomock.Any(), gomock.Any()).Return(nil, errors.New("err"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewService(storageMock, l)
			gotRes, err := s.FetchMutation(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.FetchHarga() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ServiceImpl.FetchHarga() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
