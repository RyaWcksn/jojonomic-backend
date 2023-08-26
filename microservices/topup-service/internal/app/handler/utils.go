package handler

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
