package broker

type BrokerMessage struct {
	AdminId      string  `json:"admin_id"`
	ReffId       string  `json:"reff_id"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}
