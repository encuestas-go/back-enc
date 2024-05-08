package controllers

import (
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
	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       ForumController{},
	})
}
