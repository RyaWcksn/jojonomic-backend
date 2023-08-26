package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/logger"
	"github.com/golang/mock/gomock"
)

func TestHandlerImpl_FetchHargaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := service.NewMockIService(ctrl)
	l := logger.New("", "", "")
	recorder := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantMock func()
	}{
		{
			name:    "Success",
			args:    args{w: recorder, r: httptest.NewRequest(http.MethodGet, "/api/chech-harga", nil)},
			wantErr: false,
			wantMock: func() {
				serviceMock.
					EXPECT().
					FetchHarga(gomock.Any()).
					Return(&storage.StorageEntity{
						HargaTopup:   910000,
						HargaBuyback: 820000,
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			h := NewHandler(serviceMock, l)
			if err := h.FetchHargaHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.FetchHargaHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
