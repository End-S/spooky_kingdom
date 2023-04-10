package responses

// ErrorResponse struct
type ErrorResponse struct {
	ErrMsg string `json:"error"`
}

// NewErrorResponse creates a new json error message
func NewErrorResponse(msg string) *ErrorResponse {
	err := new(ErrorResponse)
	err.ErrMsg = msg

	return err
}

// ValidationError create a new json error message for validation
func ValidationError(err error) *ErrorResponse {
	errRes := new(ErrorResponse)
	errRes.ErrMsg = err.Error()
	return errRes
}
