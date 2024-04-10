package controllers

import (
	"fmt"
	"net/http"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
)

// UserController its a struct for User
type UserController struct {
	UserRepository *repository.UserRepositoryService
}

func InitUserController() *UserController {
	repositories := repository.GetRepository()

	return &UserController{
		UserRepository: repositories.UserRespository,
	}
}

func (u *UserController) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Succesfully login  ")
}

func (u *UserController) LogOut(c echo.Context) error {
	return c.JSON(http.StatusOK, "Succesfully log out")
}

func (u *UserController) Create(c echo.Context) error {
	// get the request requirements
	user := domain.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happening trying to bind the body, err: %v", err),
		})
	}

	err = u.UserRepository.Insert(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happening trying to insert the user, the body is: %v, the error is: %v", user, err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    fmt.Sprintf("Creater %v successfully", user.Username),
	})
}

func (u *UserController) Update(c echo.Context) error {
	// que llega:
	// body del usuario
	// parametro del request
	return c.JSON(http.StatusOK, "User succesfully updated")
}

func (u *UserController) Delete(c echo.Context) error {
	// paramatro del request: id del usuario
	return c.JSON(http.StatusOK, "User succesfully deleted")
}

func (u *UserController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Here's the user selected")
}
