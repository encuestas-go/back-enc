package controllers

import (
	"fmt"
	"net/http"

	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type MapAddressController struct {
	MapRepository *repository.MapRepositoryService
}

func InitMapController(repo *repository.MapRepositoryService) *MapAddressController {
	return &MapAddressController{
		MapRepository: repo,
	}
}

func (m *MapAddressController) Get(c echo.Context) error {
	res, err := m.MapRepository.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get information of users address: %v", err),
		})
	}
	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Users Address information successfully retrieved",
		Data:       res,
	})
}
