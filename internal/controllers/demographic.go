package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DemographicStatusController struct{}

func InitDemographicController() *DemographicStatusController {
	return &DemographicStatusController{}
}

func (d *DemographicStatusController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Demographic status survey succesfully created")
}

func (d *DemographicStatusController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Demographic Status survey succesfully updated")
}

func (d *DemographicStatusController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Demographic Status survey succesfully deleted")
}

func (d *DemographicStatusController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Demographic Status survey:")
}
