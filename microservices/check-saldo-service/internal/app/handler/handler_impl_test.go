package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestHandlerImpl_FetchSaldoHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := service.NewMockIService(ctrl)
	l := logger.New("", "", "")
	recorder := httptest.NewRecorder()

	body := bytes.NewBuffer([]byte(`
{
"norek": "rek001"
}`))

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
			args:    args{w: recorder, r: httptest.NewRequest(http.MethodGet, "/api/chech-saldo", body)},
			wantErr: false,
			wantMock: func() {
				serviceMock.
					EXPECT().
					FetchSaldo(gomock.Any(), gomock.Any()).
					Return(&storage.StorageEntityRes{
						Norek:       "rek001",
						GoldBalance: 0.001,
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			h := NewHandler(serviceMock, l)
			if err := h.FetchSaldoHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.FetchHargaHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
