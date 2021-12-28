package controllers

import (
	"net/http"

	"github.com/End-S/spooky_kingdom/controllers/requests"
	"github.com/End-S/spooky_kingdom/controllers/responses"
	"github.com/End-S/spooky_kingdom/models"
	"github.com/labstack/echo/v4"
)

// AuthController struct
type AuthController struct {
	authModel *models.AuthModel
}

// NewAuthController creates a new auth controller instance
func NewAuthController(authM *models.AuthModel) *AuthController {
	return &AuthController{
		authModel: authM,
	}
}

// Login function authenticates the admin
func (ac *AuthController) Login(c echo.Context) error {
	r := new(requests.LoginReq)

	// bind json to struct
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.ValidationError(err))
	}

	jwt, err := ac.authModel.Login(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, responses.NewErrorResponse("Failed to authenticate"))
	}

	return c.JSON(http.StatusOK, responses.NewLoginResponse(jwt))
}
