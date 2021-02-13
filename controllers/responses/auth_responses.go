package responses

// LoginResponse struct
type LoginResponse struct {
	JWT string `json:"jwtToken"`
}

// NewLoginResponse creates a new login json response
func NewLoginResponse(jwt string) *LoginResponse {
	res := new(LoginResponse)
	res.JWT = jwt

	return res
}
