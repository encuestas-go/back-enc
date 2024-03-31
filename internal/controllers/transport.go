package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransportController struct{}

func InitTransportController() *TransportController {
	return &TransportController{}
}

func (t *TransportController) CreateTransportManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Transport Management survey succesfully created")
}

func (t *TransportController) UpdateTransportManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Transport Management survey succesfully updated")
}

func (t *TransportController) DeleteTransportManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Transport Management survey succesfully deleted")
}

func (t *TransportController) GetTransportManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Transport Management survey")
}
