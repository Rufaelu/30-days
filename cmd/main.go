package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/books", func(e echo.Context) error {

		return (e.JSON(http.StatusAccepted, "Ok "))
	})

	e.Start(":8080")

}
