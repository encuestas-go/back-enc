package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type SatisfactorySurveyController struct {
	SatisfactoryRepository *repository.SatisfactorySurveyRepositoryService
}

func InitSatisfactorySurveyController(repo *repository.SatisfactorySurveyRepositoryService) *SatisfactorySurveyController {
	return &SatisfactorySurveyController{
		SatisfactoryRepository: repo,
	}
}

func (s *SatisfactorySurveyController) Create(c echo.Context) error {
	idUserString := c.QueryParam("id_user")
	if idUserString == "" {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "User ID is required",
		})
	}

	idUser, err := strconv.Atoi(idUserString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid user ID: %v", err),
		})
	}

	scheduledDate := c.QueryParam("scheduled_date")
	if scheduledDate == "" {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Scheduled date is required",
		})
	}

	_, err = time.Parse("2006-01-02 15:04:05", scheduledDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid date format: %v", err),
		})
	}

	survey := domain.SatisfactorySurvey{
		IDUser:        idUser,
		ScheduledDate: scheduledDate,
	}

	err = s.SatisfactoryRepository.Insert(survey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Satisfactory survey successfully created",
	})
}

func (s *SatisfactorySurveyController) CreateLikertSurvey(c echo.Context) error {
	survey := domain.SatisfactoryLikertSurvey{}
	err := c.Bind(&survey)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	if err := s.SatisfactoryRepository.InsertLikert(survey); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Satisfactory survey successfully created",
	})
}

func (s *SatisfactorySurveyController) Get(c echo.Context) error {
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

	survey, err := s.SatisfactoryRepository.Get(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to retrieve the survey: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Satisfactory survey successfully retrieved",
		Data:       survey,
	})
}

func (s *SatisfactorySurveyController) Update(c echo.Context) error {
	survey := domain.SatisfactoryLikertSurvey{}
	if err := c.Bind(&survey); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
		})
	}

	if err := s.SatisfactoryRepository.Update(survey); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update survey",
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Survey updated",
	})
}

func (s *SatisfactorySurveyController) Delete(c echo.Context) error {
	userIDS := c.QueryParam("user_id")
	userid, err := strconv.Atoi(userIDS)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request parameters",
		})
	}

	if err := s.SatisfactoryRepository.Delete(userid); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete event",
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Survey deleted",
	})
}

func (s *SatisfactorySurveyController) GetSchedule(c echo.Context) error {
	survey, err := s.SatisfactoryRepository.GetSchedule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to retrieve the survey: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Satisfactory survey successfully retrieved",
		Data:       survey,
	})
}
