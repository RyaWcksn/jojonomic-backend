package storage

type StorageEntityRes struct {
	Norek       string  `json:"norek"`
	GoldBalance float64 `json:"saldo"`
}

type StorageEntityReq struct {
	Norek  string
	ReffId string
}
