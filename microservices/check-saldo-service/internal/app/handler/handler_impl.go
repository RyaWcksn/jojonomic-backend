package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/dto"
	"github.com/teris-io/shortid"
)

// FetchSaldoHandler implements IHandler.
func (h *HandlerImpl) FetchSaldoHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	key := shortid.MustGenerate()

	var payload dto.CheckSaldoReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.log.Errorf("ERR := %v REFF_ID := %v", err.Error(), key)
	}
	payload.ReffId = key

	h.log.Infof("Receive http request with payload := %v", payload)

	saldo, err := h.serviceImpl.FetchSaldo(ctx, &payload)
	if err != nil {
		return err
	}
	saldo.Norek = payload.Norek

	res := dto.CheckSaldoResp{
		IsError: false,
		Data:    saldo,
	}

	return ResponseJson(w, res)

}
