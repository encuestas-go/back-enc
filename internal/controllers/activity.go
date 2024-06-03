package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"

	"github.com/labstack/echo/v4"
)

type ActivityManagementController struct {
	ActivityRepository *repository.CulturalActivityRepositoryService
}

func InitActivityController(repo *repository.CulturalActivityRepositoryService) *ActivityManagementController {
	return &ActivityManagementController{
		ActivityRepository: repo,
	}
}

func (a *ActivityManagementController) Create(c echo.Context) error {
	culturalActivity := domain.CulturalActivity{}
	if err := c.Bind(&culturalActivity); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	err := a.ActivityRepository.Insert(culturalActivity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to create activity: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Activities survey successfully created",
	})
}

func (a *ActivityManagementController) Update(c echo.Context) error {
	culturalActivity := domain.CulturalActivity{}
	if err := c.Bind(&culturalActivity); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	err := a.ActivityRepository.Update(culturalActivity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to update activity: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Activities survey successfully updated",
	})
}

func (a *ActivityManagementController) Delete(c echo.Context) error {
	userIDString := c.QueryParam("user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	if err = a.ActivityRepository.Delete(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to delete activity: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Activities survey successfully deleted",
	})
}

func (a *ActivityManagementController) Get(c echo.Context) error {
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

	res, err := a.ActivityRepository.GetAllOrByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get activity: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Activities survey successfully retrieved",
		Data:       res,
	})
}

func (a *ActivityManagementController) GetCulturalActivitiesReports(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" || endDate == "" {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Start date and end date are required",
		})
	}

	res, err := a.ActivityRepository.GetCulturalActivitiesReport(startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Error generating report: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Information about cultural activities for report successfully retrieved",
		Data:       res,
	})
}
