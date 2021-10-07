package router

import (
	"net/http"

	"github.com/End-S/spooky_kingdom/config"
	"github.com/End-S/spooky_kingdom/controllers/responses"
	"github.com/End-S/spooky_kingdom/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register applies all routes to the provided router group
func Register(api *echo.Group, h *handler.Handler) {
	config := config.Get()
	admin := middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(config.JWTSecret),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSON(http.StatusForbidden, responses.NewErrorResponse("Not permitted for this resource"))
			},
		},
	)
	// AUTH
	api.POST("/login", h.AuthC.Login)
	// ARTICLES
	articles := api.Group("/articles")
	articles.GET("", h.AC.GetArticles)
	articles.POST("/update", h.AC.UpdateArticle, admin)
	// TODO change admin to API Key check
	articles.POST("/store", h.AC.StoreArticle, admin)
	articles.DELETE("/:id", h.AC.DeleteArticle, admin)
	// PUBLISHERS
	publishers := api.Group("/publishers")
	publishers.GET("", h.PC.GetPublishers)
}
