package dto

import "github.com/RyaWcksn/jojonomic-backend/check-saldo-service/internal/domain/storage"

type CheckSaldoReq struct {
	Norek  string `json:"norek"`
	ReffId string
}

type CheckSaldoResp struct {
	IsError bool                      `json:"error"`
	Data    *storage.StorageEntityRes `json:"data"`
}
