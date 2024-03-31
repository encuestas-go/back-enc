package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransportController struct{}

func InitTransportController() *TransportController {
	return &TransportController{}
}

func (t *TransportController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Transport Management survey succesfully created")
}

func (t *TransportController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Transport Management survey succesfully updated")
}

func (t *TransportController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Transport Management survey succesfully deleted")
}

func (t *TransportController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Transport Management survey")
}
