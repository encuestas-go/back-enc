package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HouseInfrastructure struct{}

func InitHouseInfrastructureController() *HouseInfrastructure {
	return &HouseInfrastructure{}
}

func (h *HouseInfrastructure) CreateHouseInfrastructure(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully created")
}

func (h *HouseInfrastructure) UpdateHouseInfrastructure(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully updated")
}

func (h *HouseInfrastructure) DeleteHouseInfrastructure(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully deleted")
}

func (h *HouseInfrastructure) GetHouseInfrastructure(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of HouseInsfrastructure survey:")
}
