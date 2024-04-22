package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type TransportController struct {
	TransportRepository *repository.TransportRespositoryService
}

func InitTransportController() *TransportController {
	repositories := repository.GetRepository()
	return &TransportController{
		TransportRepository: repositories.TransportRepository,
	}
}

func (t *TransportController) Create(c echo.Context) error {
	transportSurvey := domain.TransportManagement{}
	err := c.Bind(&transportSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = t.TransportRepository.Insert(transportSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the transport survey, the body is: %v, the error is: %v", transportSurvey, err),
		})
	}
	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Transport Management survey succesfully created",
	})
}

func (t *TransportController) Update(c echo.Context) error {
	transportSurvey := domain.TransportManagement{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}

	err = c.Bind(&transportSurvey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Unexpected error happened trying to bind the body, err: %v", err),
		})
	}
	err = t.TransportRepository.Update(transportSurvey, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Unexpected error happened trying to delete the transport survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Transport Management survey succesfully updated",
	})
}

func (t *TransportController) Delete(c echo.Context) error {
	transportSurvey := domain.TransportManagement{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}
	err = t.TransportRepository.Delete(transportSurvey, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Unexpected error happened trying to delete the transport survey: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Transport Management survey succesfully deleted",
	})
}

func (t *TransportController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Transport Management survey")
}
