package main

import (
	"fmt"
	"path/filepath"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yannis94/kanshi/internal/services/http/handlers"
)

func main() {
	fmt.Println("##### Kanshi #####")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/assets", filepath.Join("cmd", "server", "www", "public"))
	e.Renderer = handlers.NewHTMLTemplate(filepath.Join("cmd", "server", "www"))

	networkHandler := handlers.NewNetworkHandler()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "/views/index.html", map[string]interface{}{})
	})
	e.GET("/network", networkHandler.GetInfo)
	e.GET("/bandwidth", networkHandler.GetBandwidth)

	e.Logger.Fatal(e.Start(":3333"))
}
