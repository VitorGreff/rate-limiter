package main

import (
	"net/http"
	"rate-limiter/models"

	"github.com/labstack/echo/v4"
)

func main() {
	var buckets []models.Bucket
	e := echo.New()

	// e.GET("/unlimited", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Unlimited! Let's Go!")
	// })

	e.GET("/limited", func(c echo.Context) error {
		requestBucket := models.BuildBucket(c.RealIP())
    if models.IsAddrNew(buckets, requestBucket.IpAddr){
      requestBucket.ConsumeToken()
      buckets = append(buckets, requestBucket)
    }
      return c.JSON(http.StatusOK, buckets)
		})

	e.Logger.Fatal(e.Start(":8080"))
}
