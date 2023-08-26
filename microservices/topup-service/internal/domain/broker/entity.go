package broker

type BrokerMessage struct {
	ReffId       string  `json:"reff_id"`
	Type         string  `json:"type"`
	Norek        string  `json:"norek"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyBack float64 `json:"harga_buy_back"`
	GoldWeight   float64 `json:"gram"`
}
