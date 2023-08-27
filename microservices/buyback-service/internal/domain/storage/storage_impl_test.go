package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/logger"
)

func TestStorageImpl_Get(t *testing.T) {
	expect :=
		`
{
   "error": false,
   "data": {
      "norek": "rek001",
      "saldo": 0.001
   }
}
`

	svr := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, expect)
		}))

	cfg := config.Config{}
	cfg.SaldoAddr = svr.URL
	l := logger.New("dev", "dev", "debug")

	type args struct {
		ctx     context.Context
		payload *StorageEntityReq
	}
	tests := []struct {
		name    string
		args    args
		wantRes *SaldoEntity
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				payload: &StorageEntityReq{
					Norek:  "rek001",
					ReffId: "ref123",
				},
			},
			wantRes: &SaldoEntity{
				IsError: false,
				Data: Data{
					Norek:       "rek001",
					GoldBalance: 0.001,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage(cfg, l)
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
