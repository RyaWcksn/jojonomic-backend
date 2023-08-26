package dto

type TopupRequest struct {
	ReffId     string
	GoldWeight float64 `json:"gram"`
	Price      float64 `json:"harga"`
	Norek      string  `json:"norek"`
}

type TopupResponse struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id"`
}
