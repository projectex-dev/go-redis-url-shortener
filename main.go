package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to Go URL Shortener with Redis !🚀",
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
