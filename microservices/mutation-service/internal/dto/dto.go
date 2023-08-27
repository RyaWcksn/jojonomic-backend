package dto

import "github.com/RyaWcksn/jojonomic-backend/mutation-service/internal/domain/storage"

type CheckHargaResponse struct {
	IsError bool                     `json:"error"`
	Data    *[]storage.StorageEntity `json:"data"`
}

type CheckMutasiReq struct {
	Norek     string `json:"norek"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
	ReffId    string
}
