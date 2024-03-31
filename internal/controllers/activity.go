package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityController struct{}

func InitActivityController() *ActivityController {
	return &ActivityController{}
}

func (a *ActivityController) CreateActivityManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Activities survey succesfully created")
}

func (a *ActivityController) UpdateActivityManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Activities survey succesfully updated")
}

func (a *ActivityController) DeleteActivityManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Activities survey succesfully deleted")
}

func (a *ActivityController) GetActivityManagement(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Activities survey:")
}
