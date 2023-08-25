package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/input-harga-service/internal/pkgs/errors"
	"github.com/teris-io/shortid"
)

// InputHargaHandler implements IHandler.
func (h *HandlerImpl) InputHargaHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload dto.InputHargaRequest
	reffID, err := shortid.Generate()
	if err != nil {
		h.l.Errorf("Error while generating ref id", err)
		return errors.GetError("", err)
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.l.Errorf("Error wihle unmarshal payload := %v", err)
		return errors.GetError(reffID, err)
	}

	h.l.Info("NYAMPE SINI GA")

	payload.ReffID = reffID
	err = h.serviceImpl.PublishMessage(ctx, &payload)
	if err != nil {
		h.l.Errorf("Error wihle publish message := %v", err)
		return err
	}

	res := dto.InputHargaResponse{
		IsError: false,
		ReffID:  reffID,
	}

	return ResponseJson(w, res)
}
