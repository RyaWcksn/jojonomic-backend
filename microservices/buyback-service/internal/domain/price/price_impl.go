package price

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/errors"
)

// FetchPrice implements IPrice.
func (p *PriceImpl) FetchPrice(ctx context.Context, reffId string) (resp *PriceEntity, err error) {
	res, err := http.Get(p.cfg.PriceAddr + "/api/check-harga")
	if err != nil {
		p.log.Errorf("ERROR := %v , reff_id := %v", err, reffId)
		return nil, errors.GetError(reffId, err)
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		p.log.Errorf("ERROR := %v , reff_id := %v", err, reffId)
		return nil, errors.GetError(reffId, err)
	}

	p.log.Infof("Harga := %v", string(out))

	var price PriceEntity
	if err := json.Unmarshal(out, &price); err != nil {
		p.log.Errorf("ERROR := %v , reff_id := %v", err, reffId)
		return nil, errors.GetError(reffId, err)
	}

	return &price, nil
}
