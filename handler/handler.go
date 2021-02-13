package handler

import (
	"github.com/End-S/spooky_kingdom/controllers"
	"github.com/End-S/spooky_kingdom/models"
	"gorm.io/gorm"
)

type Handler struct {
	AC    *controllers.ArticleController
	AuthC *controllers.AuthController
}

func NewHandler(db *gorm.DB) *Handler {
	am := models.NewArticleModel(db)
	authm := models.NewAuthModel(db)
	return &Handler{
		AC:    controllers.NewActicleController(am),
		AuthC: controllers.NewAuthController(authm),
	}
}
