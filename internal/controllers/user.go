package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	var email, password string

	userLogin := domain.UserLogin{}
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while binding the request body: %v", err),
		})
	}

	err = u.UserRepository.Login(userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while logging in: %v", err),
		})
	}
	if userLogin.Email == email && userLogin.Password == password {
		return c.JSON(http.StatusUnauthorized, ControllerMessageResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Email or password is invalid",
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully logged in",
	})

}

func (u *UserController) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, "User succesfully logout")
}

func (u *UserController) Create(c echo.Context) error {
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
	user := domain.User{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happening trying to bind the body, err: %v", err),
		})
	}

	err = u.UserRepository.Update(user, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred when trying to update the user: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    fmt.Sprintf("User with ID %s successfully updated", userID),
	})
}

func (u *UserController) Delete(c echo.Context) error {
	user := domain.User{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}
	err = u.UserRepository.Delete(user, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while trying to delete the user: %v", err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    fmt.Sprintf("User with ID %s successfully deleted", userID),
	})
}

func (u *UserController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Here's the user selected")
}
