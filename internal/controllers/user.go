package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

type userLoginResponse struct {
	IDUser     int `json:"id_user,omitempty"`
	IDTypeUser int `json:"id_type_user,omitempty"`
}

func (u *UserController) Login(c echo.Context) error {
	//var email, password string
	userLogin := domain.UserLogin{}
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while binding the request body: %v", err),
		})
	}

	idUser, idTypeUser, err := u.UserRepository.Login(userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("An error happened when trying to login, the body is: %v, the error is: %v", userLogin, err),
		})
	}

	if idUser == 0 || idTypeUser == 0 {
		log.Printf("ID user %d ,ID Type User %d", idUser, idTypeUser)
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid email or password provided",
		})
	}

	idUserConverted := strconv.Itoa(idUser)
	idTypeUserConverted := strconv.Itoa(idTypeUser)

	cookieIDUser := new(http.Cookie)
	cookieIDUser.Name = "id_user"
	cookieIDUser.Value = idUserConverted
	cookieIDUser.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookieIDUser)

	cookieIDTypeUser := new(http.Cookie)
	cookieIDTypeUser.Name = "id_type_user"
	cookieIDTypeUser.Value = idTypeUserConverted
	cookieIDTypeUser.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookieIDTypeUser)

	return c.JSON(http.StatusOK, userLoginResponse{
		IDUser:     idUser,
		IDTypeUser: idTypeUser,
	})
}

func (u *UserController) Create(c echo.Context) error {
	user := domain.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = u.UserRepository.Insert(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the user, the body is: %v, the error is: %v", user, err),
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
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
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
