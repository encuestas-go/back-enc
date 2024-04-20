package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type SocioeconomicStatusController struct {
	SocioeconomicRepository *repository.SocioeconomicRepositoryService
}

func InitSocioeconomicController() *SocioeconomicStatusController {
	repositories := repository.GetRepository()
	return &SocioeconomicStatusController{
		SocioeconomicRepository: repositories.SocioeconomicRepository,
	}
}

func (s *SocioeconomicStatusController) Create(c echo.Context) error {
	socioeconomicSurvey := domain.SocioeconomicStatus{}
	err := c.Bind(&socioeconomicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = s.SocioeconomicRepository.Insert(socioeconomicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the socioeconomic survey, the body is: %v, the error is: %v", socioeconomicSurvey, err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Socioeconomic Status survey succesfully created",
	})
}

func (s *SocioeconomicStatusController) Update(c echo.Context) error {
	socioeconomicSurvey := domain.SocioeconomicStatus{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}

	err = c.Bind(&socioeconomicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}
	err = s.SocioeconomicRepository.Update(socioeconomicSurvey, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred when trying to update the socioeconomic survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Socioeconomic survey succesfully updated",
	})
}

func (s *SocioeconomicStatusController) Delete(c echo.Context) error {
	socioeconomicSurvey := domain.SocioeconomicStatus{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}
	err = s.SocioeconomicRepository.Delete(socioeconomicSurvey, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while trying to delete the socioeconomic survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Socioeconomic Status survey succesfully deleted",
	})
}

func (s *SocioeconomicStatusController) Get(c echo.Context) error {
	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Complete information of Socioeconomic Status survey:",
	})
}
