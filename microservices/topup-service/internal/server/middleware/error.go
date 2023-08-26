package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/RyaWcksn/jojonomic-backend/topup-service/internal/errors"
)

type ErrHandler func(http.ResponseWriter, *http.Request) error

func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := errors.ErrorForm{
				IsError: true,
				ReffId:  "",
				Message: "Panic",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		xerr := err.(*errors.ErrorForm)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(xerr)
	}

}
