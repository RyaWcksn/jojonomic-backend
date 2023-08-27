package price

import (
	"context"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/config"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/logger"
)

type IPrice interface {
	FetchPrice(ctx context.Context, refId string) (resp *PriceEntity, err error)
}

type PriceImpl struct {
	log logger.ILogger
	cfg config.Config
}

var _ IPrice = (*PriceImpl)(nil)

func NewPrice(cfg config.Config, log logger.ILogger) *PriceImpl {
	return &PriceImpl{
		log: log,
		cfg: cfg,
	}
}
