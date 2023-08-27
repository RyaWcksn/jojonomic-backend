package broker

type BrokerMessage struct {
	AdminId      string `json:"admin_id"`
	ReffId       string `json:"reff_id"`
	HargaTopup   int64  `json:"harga_topup"`
	HargaBuyback int64  `json:"harga_buyback"`
}
