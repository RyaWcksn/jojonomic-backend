package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/domain/storage"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/logger"
	"github.com/golang/mock/gomock"
)

func TestHandlerImpl_FetchHargaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := service.NewMockIService(ctrl)
	l := logger.New("", "", "")
	recorder := httptest.NewRecorder()

	body := `{
"norek": "rek123",
"start_date": 123123123,
"end_date": 123213
}`
	bodyReader := bytes.NewReader([]byte(body))

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
			args:    args{w: recorder, r: httptest.NewRequest(http.MethodPost, "/api/mutasi", bodyReader)},
			wantErr: false,
			wantMock: func() {
				serviceMock.
					EXPECT().
					FetchMutation(gomock.Any(), gomock.Any()).
					Return(&[]storage.StorageEntity{{Date: 1693092631, Type: "topup", GoldWeight: 0.001, HargaTopup: 910000, HargaBuyback: 820000, GoldBalance: 0.002}, {Date: 1693092658, Type: "topup", GoldWeight: 0.001, HargaTopup: 1000, HargaBuyback: 500, GoldBalance: 0.003}}, nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		tt.wantMock()
		t.Run(tt.name, func(t *testing.T) {
			tt.wantMock()
			h := NewHandler(serviceMock, l)
			if err := h.FetchMutationHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.FetchHargaHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
