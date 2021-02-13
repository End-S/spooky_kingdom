package main

import (
	"fmt"

	"github.com/End-S/spooky_kingdom/db"
	"github.com/End-S/spooky_kingdom/handler"
	"github.com/End-S/spooky_kingdom/router"
)

func main() {
	fmt.Println("SPOOKY KINGDOM")

	r := router.NewRouter()

	api := r.Group("/api")

	d := db.New()

	h := handler.NewHandler(d)

	router.Register(api, h)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))

	// as := models.NewArticleStore(d)

	// e.GET("/", func(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
