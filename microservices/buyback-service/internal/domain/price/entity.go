package price

type PriceEntity struct {
	IsError bool `json:"error"`
	Data    Data `json:"data"`
}

type Data struct {
	HargaTopup   int64 `json:"harga_topup"`
	HargaBuyback int64 `json:"harga_buyback"`
}
