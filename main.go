package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
  e := echo.New()

  e.GET("/unlimited", func(c echo.Context) error {
    return c.String(http.StatusOK, "Unlimited! Let's Go!")
  })
  
  e.GET("/limited", func(c echo.Context) error {
    return c.String(http.StatusOK, "Limited, don't over use me!")
  })


  e.Logger.Fatal(e.Start(":8080"))
}
