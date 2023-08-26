package handler

import (
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/app/service"
	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/logger"
)

type IHandler interface {
	TopupHandler(w http.ResponseWriter, r *http.Request) error
}

type HandlerImpl struct {
	serviceImpl service.IService
	l           logger.ILogger
}

var _ IHandler = (*HandlerImpl)(nil)

func NewHandler(s service.IService, l logger.ILogger) *HandlerImpl {
	return &HandlerImpl{
		serviceImpl: s,
		l:           l,
	}
}
