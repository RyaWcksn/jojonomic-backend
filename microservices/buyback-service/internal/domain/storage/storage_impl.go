package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/buyback-service/internal/pkgs/errors"
)

// Get implements IStorage.
func (s *StorageImpl) Get(ctx context.Context, payload *StorageEntityReq) (res *SaldoEntity, err error) {
	body := fmt.Sprintf(`{"norek": "%s"}`, payload.Norek)
	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)

	httpRes, err := http.Post(s.cfg.SaldoAddr+"/api/check-saldo", "application/json", bodyReader)
	if err != nil {
		s.log.Errorf("ERROR := %v , reff_id := %v", err, payload.ReffId)
		return nil, errors.GetError(payload.ReffId, err)
	}
	defer httpRes.Body.Close()

	out, err := io.ReadAll(httpRes.Body)
	if err != nil {
		s.log.Errorf("ERROR := %v , reff_id := %v", err.Error(), payload.ReffId)
		return nil, errors.GetError(payload.ReffId, err)
	}
	fmt.Println("SALDO := ", string(out))

	s.log.Infof("SALDO := %v", string(out))

	var saldo SaldoEntity
	if err := json.Unmarshal(out, &saldo); err != nil {
		s.log.Errorf("ERROR := %v , reff_id := %v", err, payload.ReffId)
		return nil, errors.GetError(payload.ReffId, err)
	}

	return &saldo, nil

}
