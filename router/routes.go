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
	// JWT Auth
	admin := middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(config.JWTSecret),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return echo.NewHTTPError(http.StatusForbidden, responses.NewErrorResponse("Not permitted for this resource"))
			},
		},
	)
	// API Key Auth
	key := middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == config.APIKey, nil
	})
	// AUTH
	api.POST("/login", h.AuthC.Login)
	// ARTICLES
	articles := api.Group("/articles")
	articles.GET("", h.AC.GetArticles)
	articles.PATCH("/:id", h.AC.UpdateArticle, admin)
	articles.POST("", h.AC.StoreArticle, key)
	articles.DELETE("/:id", h.AC.DeleteArticle, admin)
	articles.GET("/dates", h.AC.ArticleDateSpan)
	// PUBLISHERS
	publishers := api.Group("/publishers")
	publishers.GET("", h.PC.GetPublishers)
	publishers.POST("", h.PC.CreatePublisher, key)
}
