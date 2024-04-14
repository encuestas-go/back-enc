package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityManagementController struct{}

func InitActivityController() *ActivityManagementController {
	return &ActivityManagementController{}
}

func (a *ActivityManagementController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Activities survey succesfully created")
}

func (a *ActivityManagementController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Activities survey succesfully updated")
}

func (a *ActivityManagementController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Activities survey succesfully deleted")
}

func (a *ActivityManagementController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Activities survey:")
}
