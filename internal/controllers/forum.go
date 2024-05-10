package controllers

import (
	"fmt"
	"net/http"

	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

type ForumController struct {
	ForumRepository *repository.ForumRepositoryService
}

func InitForumController(repo *repository.ForumRepositoryService) *ForumController {
	return &ForumController{
		ForumRepository: repo,
	}
}

func (f *ForumController) Get(c echo.Context) error {

	res, err := f.ForumRepository.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get information of forum: %v", err),
		})
	}
	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Forum information successfully retrieved",
		Data:       res,
	})
}
