package storage

type StorageEntity struct {
	Date         int64   `json:"date"`
	Type         string  `json:"type"`
	GoldWeight   float64 `json:"gram"`
	HargaTopup   int64   `json:"harga_topup"`
	HargaBuyback int64   `json:"harga_buyback"`
	GoldBalance  float64 `json:"saldo"`
}

type StorageRequest struct {
	From   int64
	To     int64
	ReffId string
	Norek  string
}
