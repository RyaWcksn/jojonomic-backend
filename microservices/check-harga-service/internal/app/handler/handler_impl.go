package handler

import (
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/dto"
)

// FetchHargaHandler implements IHandler.
func (h *HandlerImpl) FetchHargaHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	harga, err := h.serviceImpl.FetchHarga(ctx)
	if err != nil {
		return err
	}

	res := dto.CheckHargaResponse{
		IsError: false,
		Data:    harga,
	}

	return ResponseJson(w, res)

}
