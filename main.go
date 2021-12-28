package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/End-S/spooky_kingdom/db"
	"github.com/End-S/spooky_kingdom/db/migration"
	"github.com/End-S/spooky_kingdom/handler"
	"github.com/End-S/spooky_kingdom/router"
	"github.com/labstack/echo/v4"
)

//go:embed ui/dist
var embeddedFiles embed.FS

func main() {
	var banner = ` 
 ____                    _            _  ___                 _                 
/ ___| _ __   ___   ___ | | ___   _  | |/ (_)_ __   __ _  __| | ___  _ __ ___  
\___ \| '_ \ / _ \ / _ \| |/ / | | | | ' /| | '_ \ / _  |/ _  |/ _ \|  _   _ \ 
 ___) | |_) | (_) | (_) |   <| |_| | | . \| | | | | (_| | (_| | (_) | | | | | |
|____/| .__/ \___/ \___/|_|\_\\__, | |_|\_\_|_| |_|\__, |\__,_|\___/|_| |_| |_|
      |_|                     |___/                |___/                       
`
	fmt.Println(banner)

	migration.MigrateUp()

	r := router.NewRouter()

	api := r.Group("/api")

	d := db.New()

	h := handler.NewHandler(d)

	router.Register(api, h)

	// serve static files
	r.GET("/*", echo.WrapHandler(clientHandler()))

	r.Logger.Fatal(r.Start("localhost:8585"))
}

func clientHandler() http.Handler {
	staticClient, err := fs.Sub(embeddedFiles, "ui/dist")
	if err != nil {
		panic(err)
	}
	fmt.Println("🫖 Serving embedded static files")
	return http.FileServer(http.FS(staticClient))
}
