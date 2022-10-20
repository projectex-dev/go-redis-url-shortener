package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-redis-url-shortener/handler"
	"go-redis-url-shortener/store"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to Go URL Shortener with Redis !🚀",
		})
	})

	e.POST("/encode", func(c echo.Context) error {
		return handler.CreateShortURL(c)
	})

	e.GET("/decode/:short-url", func(c echo.Context) error {
		return handler.ReturnLongURL(c)
	})

	// Store initialization happens here
	store.InitializeStore()

	e.Logger.Fatal(e.Start(":1323"))
}
