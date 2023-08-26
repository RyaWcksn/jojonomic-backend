package price

type PriceEntity struct {
	IsError bool `json:"error"`
	Data    Data `json:"data"`
}

type Data struct {
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}
