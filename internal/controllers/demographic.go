package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DemographicController struct{}

func InitDemographicController() *DemographicController {
	return &DemographicController{}
}

func (d *DemographicController) CreateDemographicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Demographic status survey succesfully created")
}

func (d *DemographicController) UpdateDemographicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Demographic Status survey succesfully updated")
}

func (d *DemographicController) DeleteDemographicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Demographic Status survey succesfully deleted")
}

func (d *DemographicController) GetDemographicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Demographic Status survey:")
}
