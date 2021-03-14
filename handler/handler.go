package handler

import (
	"github.com/End-S/spooky_kingdom/controllers"
	"github.com/End-S/spooky_kingdom/models"
	"gorm.io/gorm"
)

// Handler struct
type Handler struct {
	AC    *controllers.ArticleController
	AuthC *controllers.AuthController
	PC    *controllers.PublisherController
}

// NewHandler creates a new Handler
func NewHandler(db *gorm.DB) *Handler {
	am := models.NewArticleModel(db)
	authm := models.NewAuthModel(db)
	pm := models.NewPublisherModel(db)
	return &Handler{
		AC:    controllers.NewActicleController(am),
		AuthC: controllers.NewAuthController(authm),
		PC:    controllers.NewPublisherController(pm),
	}
}
