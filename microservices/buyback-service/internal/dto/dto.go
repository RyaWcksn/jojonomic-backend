package dto

type BuybackRequest struct {
	ReffId     string
	GoldWeight float64 `json:"gram"`
	Price      int64   `json:"harga"`
	Norek      string  `json:"norek"`
}

type BuybackResponse struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id"`
}
