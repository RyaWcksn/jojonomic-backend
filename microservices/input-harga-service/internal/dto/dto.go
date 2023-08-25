package dto

type InputHargaRequest struct {
	AdminId      string  `json:"admin_id"`
	ReffID       string  `json:"reff_id"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}

type InputHargaResponse struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id"`
}
