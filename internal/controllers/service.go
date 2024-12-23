package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"

	"github.com/labstack/echo/v4"
)

type ServiceManagementController struct {
	ServiceRepository *repository.ServicesRepositoryService
}

func InitServiceManagementController(repo *repository.ServicesRepositoryService) *ServiceManagementController {
	return &ServiceManagementController{
		ServiceRepository: repo,
	}
}

func (sm *ServiceManagementController) Create(c echo.Context) error {
	service := domain.Services{}
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	if err := sm.ServiceRepository.Insert(service); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error inserting service: %s", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "service created",
	})
}

func (sm *ServiceManagementController) Update(c echo.Context) error {
	service := domain.Services{}
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	if err := sm.ServiceRepository.Update(service); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error updating service: %s", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "service updated",
	})
}

func (sm *ServiceManagementController) Delete(c echo.Context) error {
	userIDString := c.QueryParam("user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	if err = sm.ServiceRepository.Delete(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error deleting service: %s", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "service deleted",
	})
}

func (sm *ServiceManagementController) Get(c echo.Context) error {
	userIDString := c.QueryParam("user_id")
	if userIDString == "" {
		userIDString = "0"
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	services, err := sm.ServiceRepository.GetAllOrByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("error to retrived service: %s", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "service retrieved",
		Data:       services,
	})
}

func (sm *ServiceManagementController) GetAllInternetProvidersReport(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" || endDate == "" {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Start date and end date are required",
		})
	}

	res, err := sm.ServiceRepository.GetInternetProviderReport(startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Error generating report: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Information of internet providers report successfully retrieved",
		Data:       res,
	})
}
