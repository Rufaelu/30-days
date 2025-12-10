package routes

import (
	"30days/models"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()

	e.GET("/create", func(e echo.Context) error {
		conn, err := models.Connect()
		if err != nil {
			return (e.JSON(http.StatusBadRequest, err.Error()))

		}
		err = models.CreateDB("Mine", conn)
		if err != nil {
			return (e.JSON(http.StatusBadRequest, err.Error()))

		}
		defer conn.Close(context.Background())

		return (e.JSON(http.StatusAccepted, "Done"))
	})

	return e

}
