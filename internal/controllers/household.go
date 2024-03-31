package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HouseInfrastructureController struct{}

func InitHouseInfrastructureController() *HouseInfrastructureController {
	return &HouseInfrastructureController{}
}

func (h *HouseInfrastructureController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully created")
}

func (h *HouseInfrastructureController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully updated")
}

func (h *HouseInfrastructureController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully deleted")
}

func (h *HouseInfrastructureController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of HouseInsfrastructure survey:")
}
