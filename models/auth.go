package models

import (
	"errors"
	"fmt"

	"github.com/End-S/spooky_kingdom/config"
	"github.com/End-S/spooky_kingdom/controllers/requests"
	"github.com/End-S/spooky_kingdom/utils"
	"gorm.io/gorm"
)

// AuthModel struct
type AuthModel struct {
	db *gorm.DB
}

// NewAuthModel creates a new auth model instance
func NewAuthModel(db *gorm.DB) *AuthModel {
	return &AuthModel{
		db: db,
	}
}

// Login function returns a valid JWT
func (am *AuthModel) Login(req *requests.LoginReq) (string, error) {
	fmt.Println(req)
	config := config.Get()

	if req.Password != config.Admin.Password ||
		req.Username != config.Admin.Username {
		return "", errors.New("Invalid login")
	}

	return utils.GenerateJWT(), nil
}
