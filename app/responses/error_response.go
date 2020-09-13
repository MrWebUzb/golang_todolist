package responses

import (
	"net/http"
)

// ErrorResponse for http errors
type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse ...
func NewErrorResponse(code int) *ErrorResponse {
	return &ErrorResponse{
		Ok:      false,
		Code:    code,
		Message: http.StatusText(code),
	}
}
