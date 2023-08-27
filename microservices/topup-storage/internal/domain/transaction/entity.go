package transaction

type TransactionEntity struct {
	ReffId       string  `json:"reff_id"`
	Type         string  `json:"type"`
	Norek        string  `json:"norek"`
	HargaTopup   int64   `json:"harga_topup"`
	HargaBuyBack int64   `json:"harga_buy_back"`
	GoldWeight   float64 `json:"gram"`
	GoldBalance  float64 `json:"saldo"`
	CreatedAt    int
}
