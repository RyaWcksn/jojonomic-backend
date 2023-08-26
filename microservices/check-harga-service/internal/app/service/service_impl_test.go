package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/logger"
	"github.com/golang/mock/gomock"
)

func TestServiceImpl_FetchHarga(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := storage.NewMockIStorage(ctrl)
	l := logger.New("", "", "")
	ctx := context.Background()

	tests := []struct {
		name     string
		wantRes  *storage.StorageEntity
		wantErr  bool
		wantMock func()
	}{
		{
			name: "Success",
			wantRes: &storage.StorageEntity{
				HargaTopup:   910000,
				HargaBuyback: 820000,
			},
			wantErr: false,
			wantMock: func() {
				storageMock.
					EXPECT().
					FetchHarga(gomock.Any()).
					Return(&storage.StorageEntity{
						HargaTopup:   910000,
						HargaBuyback: 820000,
					}, nil)
			},
		},
		{
			name:    "Failed",
			wantRes: nil,
			wantErr: true,
			wantMock: func() {
				storageMock.
					EXPECT().
					FetchHarga(gomock.Any()).
					Return(
						nil, errors.New("err"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewService(storageMock, l)
			gotRes, err := s.FetchHarga(ctx)
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
