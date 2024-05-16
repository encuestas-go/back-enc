package controllers

import (
	"net/http"
	"strconv"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"

	"github.com/labstack/echo/v4"
)

type EventManagementController struct {
	EventRepository *repository.EventRepositoryService
}

func InitEventManagementController(repo *repository.EventRepositoryService) *EventManagementController {
	return &EventManagementController{
		EventRepository: repo,
	}
}

func (em *EventManagementController) Create(c echo.Context) error {
	event := domain.Event{}
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
		})
	}

	if err := em.EventRepository.CreateEvent(event); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create event",
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Event created",
	})
}

func (em *EventManagementController) Update(c echo.Context) error {
	event := domain.Event{}
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
		})
	}

	if err := em.EventRepository.Update(event); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update event",
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Event updated",
	})
}

func (em *EventManagementController) Delete(c echo.Context) error {
	userIDS := c.QueryParam("user_id")
	userid, err := strconv.Atoi(userIDS)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request parameters",
		})
	}

	if err := em.EventRepository.DeleteEvent(userid); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to delete event",
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Event deleted",
	})
}

func (em *EventManagementController) Get(c echo.Context) error {
	events, err := em.EventRepository.GetEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get events",
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Events retrieved",
		Data:       events,
	})
}
