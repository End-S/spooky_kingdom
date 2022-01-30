package controllers

import (
	"fmt"
	"net/http"

	"github.com/End-S/spooky_kingdom/controllers/requests"
	"github.com/End-S/spooky_kingdom/controllers/responses"
	"github.com/End-S/spooky_kingdom/models"
	"github.com/End-S/spooky_kingdom/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ArticleController struct
type ArticleController struct {
	articleModel *models.ArticleModel
}

// NewArticleController creates a new article controller instance
func NewArticleController(am *models.ArticleModel) *ArticleController {
	return &ArticleController{
		articleModel: am,
	}
}

// GetArticles function handles a request for articles
func (ac *ArticleController) GetArticles(c echo.Context) error {
	r := new(requests.GetArticlesReq)

	// bind query params to struct
	if err := c.Bind(r); err != nil {
		return err
	}

	// validate request
	if err := c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.ValidationError(err))
	}

	// if user want to retrieve articles that are pending review check they are admin
	if r.AssessPending == "true" {
		if err := utils.VerifyAuthHeader(c, "admin"); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, responses.NewErrorResponse("Not permitted for this resource"))
		}
	} else {
		r.AssessPending = "false"
	}

	articles, count, err := ac.articleModel.List(r, r.AssessPending == "true")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to retrieve articles"))
	}

	return c.JSON(http.StatusOK, responses.NewListArticlesResponse(articles, count))
}

// UpdateArticle function handles the updating of an article
func (ac *ArticleController) UpdateArticle(c echo.Context) error {
	r := new(requests.UpdateArticleReq)

	// bind json to struct
	if err := c.Bind(r); err != nil {
		return err
	}

	// validate request
	if err := c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.ValidationError(err))
	}

	article, err := ac.articleModel.Update(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to update article"))
	}

	return c.JSON(http.StatusOK, responses.NewArticleResponse(*article))
}

// StoreArticle function handles the storing of an article
func (ac *ArticleController) StoreArticle(c echo.Context) error {
	r := new(requests.StoreArticleReq)

	// bind json to struct
	if err := c.Bind(r); err != nil {
		fmt.Println(err)
		return err
	}

	// validate request
	if err := c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.ValidationError(err))
	}

	article, err := ac.articleModel.Store(r)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.NewErrorResponse("Server error, unable to store article"))
	}

	return c.JSON(http.StatusOK, responses.NewArticleResponse(*article))
}

// DeleteArticle function handles a delete request for an article
func (ac *ArticleController) DeleteArticle(c echo.Context) error {
	// extract param from url
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			responses.NewErrorResponse("Unrecognized UUID"))
	}

	count, err := ac.articleModel.Delete(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			responses.NewErrorResponse("Server error, unable to delete article"))
	}

	if count <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest,
			responses.NewErrorResponse("Article does not exist"))
	}

	return c.JSON(http.StatusOK, responses.NewDeleteResponse(true))
}

// ArticleDateSpan function handles an article date span request
func (ac *ArticleController) ArticleDateSpan(c echo.Context) error {
	dateSpan, err := ac.articleModel.DateSpan()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			responses.NewErrorResponse("Server error, unable to get article date span"))
	}

	return c.JSON(http.StatusOK, responses.NewArticleDateSpanResponse(*dateSpan))
}

// To improve pagination performance we should have indexes for fields, which we are using in ORDER BY
// index date pub, title ASC/DSC and type
