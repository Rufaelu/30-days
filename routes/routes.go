package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	e.GET("/books", func(e echo.Context) error {

		return (e.JSON(http.StatusAccepted, "Ok "))
	})

	return e

}
