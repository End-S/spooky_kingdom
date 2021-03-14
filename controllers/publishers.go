package controllers

import (
	"net/http"

	"github.com/End-S/spooky_kingdom/controllers/responses"
	"github.com/End-S/spooky_kingdom/models"
	"github.com/labstack/echo/v4"
)

//PublisherController struct
type PublisherController struct {
	publisherModel *models.PublisherModel
}

// NewPublisherController creates a new publisher controller instance
func NewPublisherController(pm *models.PublisherModel) *PublisherController {
	return &PublisherController{
		publisherModel: pm,
	}
}

// GetPublishers function handles a request for publishers
func (ac *PublisherController) GetPublishers(c echo.Context) error {
	publishers, err := ac.publisherModel.List()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to retrieve publishers"))
	}

	return c.JSON(http.StatusOK, responses.NewListPublishersResponse(publishers))
}

// GET all publisers

// ADMIN POST create new publiser

// ADMIN POST publishers update
