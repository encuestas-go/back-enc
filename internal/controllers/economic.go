package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EconomicStatusController struct{}

func InitEconomicController() *EconomicStatusController {
	return &EconomicStatusController{}
}

func (e *EconomicStatusController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Economic Status survey succesfully created")
}

func (e *EconomicStatusController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Economic Status survey succesfully updated")
}

func (e *EconomicStatusController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Economic Status survey succesfully deleted")
}

func (e *EconomicStatusController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Economic Status survey :")
}
