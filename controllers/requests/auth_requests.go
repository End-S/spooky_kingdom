package requests

import "errors"

// LoginReq struct for a Login request
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginReq) Validate() error {
	if req.Username == "" {
		return errors.New("Username is required")
	}
	if req.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}
