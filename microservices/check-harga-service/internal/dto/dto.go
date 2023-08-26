package dto

import "github.com/RyaWcksn/jojonomic-backend/check-harga-service/internal/domain/storage"

type CheckHargaResponse struct {
	IsError bool                   `json:"error"`
	Data    *storage.StorageEntity `json:"data"`
}
