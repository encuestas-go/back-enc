package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type EconomicStatusController struct {
	EconomicRepository *repository.EconomicRepositoryService
}

func InitEconomicController() *EconomicStatusController {
	repositories := repository.GetRepository()
	return &EconomicStatusController{
		EconomicRepository: repositories.EconomicRepository,
	}
}

func (e *EconomicStatusController) Create(c echo.Context) error {
	economicSurvey := domain.EconomicStatus{}
	err := c.Bind(&economicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = e.EconomicRepository.Insert(economicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the economic survey, the body is: %v, the error is: %v", economicSurvey, err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Economic Status survey succesfully created",
	})
}

func (e *EconomicStatusController) Update(c echo.Context) error {
	economicSurvey := domain.EconomicStatus{}
	err := c.Bind(&economicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}
	err = e.EconomicRepository.Update(economicSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred when trying to update the economic survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Economic survey succesfully updated",
	})
}

func (e *EconomicStatusController) Delete(c echo.Context) error {
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}

	err = e.EconomicRepository.Delete(userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while trying to delete the economic survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Economic Status survey succesfully deleted",
	})
}

func (e *EconomicStatusController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Economic Status survey :")
}
