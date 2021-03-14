package controllers

import (
	"net/http"

	"github.com/End-S/spooky_kingdom/controllers/requests"
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
func (pc *PublisherController) GetPublishers(c echo.Context) error {
	publishers, err := pc.publisherModel.List()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to retrieve publishers"))
	}

	return c.JSON(http.StatusOK, responses.NewListPublishersResponse(publishers))
}

// GET all publisers

// ADMIN POST create new publisher
func (pc *PublisherController) CreatePublisher(c echo.Context) error {
	r := new(requests.CreatePublisherReq)

	// bind query params to struct
	if err := c.Bind(r); err != nil {
		return err
	}

	// validate request
	if err := c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.ValidationError(err))
	}

	publisher, err := pc.publisherModel.Create(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to create Publisher"))
	}

	return c.JSON(http.StatusOK, responses.NewPublisherResponse(*publisher))

}

// ADMIN POST publishers update
