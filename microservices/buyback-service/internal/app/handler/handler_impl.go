package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/dto"
	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/errors"
	"github.com/teris-io/shortid"
)

// BuybackHandler implements IHandler.
func (h *HandlerImpl) BuybackHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload dto.BuybackRequest
	reffID, err := shortid.Generate()
	if err != nil {
		h.l.Errorf("Error while generating ref id", err)
		return errors.GetError("", err)
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.l.Errorf("Error wihle unmarshal payload := %v", err)
		return errors.GetError(reffID, err)
	}

	payload.ReffId = reffID
	err = h.serviceImpl.PublishMessage(ctx, &payload)
	if err != nil {
		h.l.Errorf("Error wihle publish message := %v", err)
		return err
	}

	res := dto.BuybackResponse{
		IsError: false,
		ReffID:  reffID,
	}

	return ResponseJson(w, res)
}
