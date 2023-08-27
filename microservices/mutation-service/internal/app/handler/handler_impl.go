package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/pkgs/errors"
	"github.com/teris-io/shortid"
)

// FetchMutationHandler implements IHandler.
func (h *HandlerImpl) FetchMutationHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	key := shortid.MustGenerate()
	var payload dto.CheckMutasiReq

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.log.Errorf("Err := %v REFF_ID %v", err, key)
		return errors.GetError(key, err)
	}

	saldo, err := h.serviceImpl.FetchMutation(ctx, &payload)
	if err != nil {
		return err
	}

	res := dto.CheckHargaResponse{
		IsError: false,
		Data:    saldo,
	}

	return ResponseJson(w, res)

}
