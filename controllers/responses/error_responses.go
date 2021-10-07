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

// ValidationError create a new json error message for validation
// func ValidationError(vErr error, i interface{}) *ErrorResponse {
// 	// cast err to ValidationErrors
// 	errors := vErr.(validator.ValidationErrors)
// 	// reflect the interface so we can read the err tag
// 	r := reflect.TypeOf(i)

// 	msg := ""
// 	for _, err := range errors {
// 		f, _ := r.FieldByName(err.StructField())
// 		tagHelp, found := f.Tag.Lookup(err.ActualTag())
// 		// TODO refactor string concat
// 		if found {
// 			msg += tagHelp + ". "
// 		} else {
// 			msg += err.Field() + " " + err.ActualTag() + " " + err.Param()
// 		}

// 	}

// 	err := new(ErrorResponse)
// 	err.ErrMsg = msg

// 	return err
// }
