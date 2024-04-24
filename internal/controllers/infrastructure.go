package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type HouseInfrastructureController struct {
	InfrastructureRepository *repository.HouseInfrastructureRepositoryService
}

func InitHouseInfrastructureController() *HouseInfrastructureController {
	repositories := *repository.GetRepository()
	return &HouseInfrastructureController{
		InfrastructureRepository: repositories.InfrastructureRepository,
	}
}

func (h *HouseInfrastructureController) Create(c echo.Context) error {
	householdSurvey := domain.HouseholdInfrastructure{}
	err := c.Bind(&householdSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = h.InfrastructureRepository.Insert(householdSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the household infrastructure survey, the body is: %v, the error is: %v", householdSurvey, err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Household Infrastructure survey succesfully created",
	})
}

func (h *HouseInfrastructureController) Update(c echo.Context) error {
	householdSurvey := domain.HouseholdInfrastructure{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}

	err = c.Bind(&householdSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}
	err = h.InfrastructureRepository.Update(householdSurvey, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred when trying to update the household infrastructure survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Household Infrastructure survey succesfully updated",
	})
}

func (h *HouseInfrastructureController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Household Infrastructure survey succesfully deleted")
}

func (h *HouseInfrastructureController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of HouseInsfrastructure survey:")
}