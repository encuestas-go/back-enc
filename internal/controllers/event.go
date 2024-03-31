package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventManagementController struct{}

func InitEventManagementController() *EventManagementController {
	return &EventManagementController{}
}

func (em *EventManagementController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Event survey succesfully created")
}

func (em *EventManagementController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Event survey succesfully updated")
}

func (em *EventManagementController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Event survey succesfully deleted")
}

func (em *EventManagementController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Event survey:")
}
