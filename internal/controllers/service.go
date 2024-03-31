package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServiceManagementController struct{}

func InitServiceManagementController() *ServiceManagementController {
	return &ServiceManagementController{}
}

func (sm *ServiceManagementController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Service Management survey succesfully created")
}

func (sm *ServiceManagementController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Service Management survey succesfully updated")
}

func (sm *ServiceManagementController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Service Management survey succesfully deleted")
}

func (sm *ServiceManagementController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Service Management survey:")
}
