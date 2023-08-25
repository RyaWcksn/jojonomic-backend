package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestHandlerImpl_InputHargaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := service.NewMockIService(ctrl)
	l := logger.Init("", "", "")

	recorder := httptest.NewRecorder()
	successRequestBody := bytes.NewBuffer([]byte(`
{
"admin_id": "a001",
"harga_topup": 910000,
"harga_buyback": 820000
}`))

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantMock func()
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				w: recorder,
				r: httptest.NewRequest(http.MethodPost, "/api/input-harga/", successRequestBody),
			},
			wantMock: func() {
				serviceMock.EXPECT().PublishMessage(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			h := NewHandler(serviceMock, l)
			if err := h.InputHargaHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.InputHargaHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
