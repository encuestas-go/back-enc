package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EconomicController struct{}

func InitEconomicController() *EconomicController {
	return &EconomicController{}
}

func (e *EconomicController) CreateEconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Economic Status survey succesfully created")
}

func (e *EconomicController) UpdateEconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Economic Status survey succesfully updated")
}

func (e *EconomicController) DeleteEconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Economic Status survey succesfully deleted")
}

func (e *EconomicController) GetEconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Economic Status survey :")
}
