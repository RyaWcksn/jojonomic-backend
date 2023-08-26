package handler

import (
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/logger"
)

type IHandler interface {
	FetchHargaHandler(w http.ResponseWriter, r *http.Request) error
}

type HandlerImpl struct {
	serviceImpl service.IService
	log         logger.ILogger
}

var _ IHandler = (*HandlerImpl)(nil)

func NewHandler(s service.IService, l logger.ILogger) *HandlerImpl {
	return &HandlerImpl{
		serviceImpl: s,
		log:         l,
	}
}
