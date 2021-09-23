package errors

import "net/http"

type ErrResponse struct {
	Message string
	Status  int
	Error   string
}

func NewInternalServerError(message string) *ErrResponse {
	return &ErrResponse{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewBadRequestError(message string) *ErrResponse {
	return &ErrResponse{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request_error",
	}
}
