package controllers

import (
	"fmt"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ForumController struct {
	ForumRepository *repository.ForumRepositoryService
}

func InitForumController(repo *repository.ForumRepositoryService) *ForumController {
	return &ForumController{
		ForumRepository: repo,
	}
}

func (f *ForumController) CreateQuestion(c echo.Context) error {
	question := domain.Question{}
	err := c.Bind(&question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body of the question, err: %v", err),
		})
	}

	questionID, err := f.ForumRepository.InsertQuestion(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the question, the body is: %v, the error is: %v", question, err),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status_code": http.StatusCreated,
		"message":     "Question successfully created",
		"question_id": questionID,
	})
}

func (f *ForumController) CreateAnswer(c echo.Context) error {
	answer := domain.Answer{}
	err := c.Bind(&answer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body of the answer, err: %v", err),
		})
	}

	err = f.ForumRepository.InsertAnswer(answer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the answer, the body is: %v, the error is: %v", answer, err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Answer successfully created",
	})
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
