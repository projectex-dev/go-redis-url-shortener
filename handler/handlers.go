package handler

import (
	"github.com/labstack/echo/v4"
	"go-redis-url-shortener/shortener"
	"go-redis-url-shortener/store"
)

const host = "http://localhost:1323/"

// URLCreationRequest is request model definition
type URLCreationRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c echo.Context) error {
	cr := new(URLCreationRequest)
	if err := c.Bind(cr); err != nil {
		return err
	}

	shortUrl := shortener.GenerateShortURL(cr.LongURL, cr.UserId)
	store.SaveURLInRedis(shortUrl, cr.LongURL)

	return c.JSON(200, map[string]interface{}{
		"short_url": host + shortUrl,
	})
}

func ReturnLongURL(c echo.Context) error {
	shortUrl := c.Param("short-url")
	initialUrl := store.RetrieveInitialURLFromRedis(shortUrl)
	return c.JSON(200, map[string]interface{}{
		"short_url": host + shortUrl,
		"long_url":  initialUrl,
	})
}
