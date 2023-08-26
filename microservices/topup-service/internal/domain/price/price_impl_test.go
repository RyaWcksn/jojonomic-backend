package price

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/logger"
)

func TestPriceImpl_FetchPrice(t *testing.T) {

	expect := `
{
"error": false,
"data": {
   "harga_topup": 910000,
   "harga_buyback": 820000
}
}
`

	svr := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, expect)
		}))

	cfg := config.Config{}
	cfg.PriceAddr = svr.URL

	l := logger.New("", "", "")

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		wantResp *PriceEntity
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
			},
			wantResp: &PriceEntity{
				IsError: false,
				Data: Data{
					HargaTopup:   910000,
					HargaBuyback: 820000,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPrice(cfg, l)
			gotResp, err := p.FetchPrice(tt.args.ctx, "res_123")
			if (err != nil) != tt.wantErr {
				t.Errorf("PriceImpl.FetchPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PriceImpl.FetchPrice() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
