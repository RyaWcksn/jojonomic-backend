package errors

import "fmt"

type IError interface {
	Error() string
}

type ErrorForm struct {
	IsError bool   `json:"error"`
	ReffId  string `json:"reff_id"`
	Message string `json:"message"`
}

func (o ErrorForm) Error() string {
	return fmt.Sprintf("IsError = %v Reff_id := %v errors = %v", o.IsError, o.ReffId, o.Message)
}

// GetError code and message then return.
func GetError(reffId string, errActual error) *ErrorForm {
	return &ErrorForm{
		IsError: true,
		ReffId:  reffId,
		Message: errActual.Error(),
	}
}
