package service

import (
	"context"
	"errors"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/domain/broker"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/dto"
	rr "github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/errors"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestServiceImpl_PublishMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	brokerMock := broker.NewMockIBroker(ctrl)
	l := logger.Init("", "", "")

	type args struct {
		ctx     context.Context
		payload *dto.InputHargaRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantMock func()
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: &dto.InputHargaRequest{
					AdminId:      "a001",
					HargaTopup:   91000,
					HargaBuyback: 82000,
				},
			},
			wantErr: false,
			wantMock: func() {
				brokerMock.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed",
			args: args{
				ctx:     context.Background(),
				payload: &dto.InputHargaRequest{},
			},
			wantErr: true,
			wantMock: func() {
				brokerMock.
					EXPECT().
					Publish(gomock.Any(), gomock.Any()).
					Return(rr.GetError("123123", errors.New("Error")))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			s := NewService(brokerMock, l)
			if err := s.PublishMessage(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.PublishMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
