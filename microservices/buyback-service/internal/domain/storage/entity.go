package storage

type StorageEntityRes struct {
	GoldBalance float64
}

type StorageEntityReq struct {
	Norek  string
	ReffId string
}

type SaldoEntity struct {
	IsError bool `json:"error"`
	Data    Data `json:"data"`
}

type Data struct {
	Norek       string  `json:"norek"`
	GoldBalance float64 `json:"saldo"`
}
