package main

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yannis94/kanshi/internal/services/http/handlers"
)

func main() {
	fmt.Println("##### Kanshi #####")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	networkHandler := handlers.NewNetworkHandler()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello friend"})
	})
	e.GET("/network", networkHandler.GetInfo)

	e.Logger.Fatal(e.Start(":3333"))
}
