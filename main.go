package main

import (
	"net/http"
	"rate-limiter/models"

	"github.com/labstack/echo/v4"
)

func main() {
	bucket := models.BuildBucket("vitao", 4.0)
	e := echo.New()

	e.GET("/unlimited", func(c echo.Context) error {
		return c.String(http.StatusOK, "Unlimited! Let's Go!")
	})

	e.GET("/limited", func(c echo.Context) error {
		err := bucket.TakeToken()
		if err != nil {
			return c.JSON(http.StatusTooManyRequests, err.Error())
		}
		return c.JSON(http.StatusOK, bucket)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
