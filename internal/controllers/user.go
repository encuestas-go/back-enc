package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserController its a struct for User
type UserController struct{}

func InitUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(c echo.Context) error {
	return c.JSON(200, "User succesfully created ")
}

func (u *UserController) UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "User succesfully updated")
}

func (u *UserController) DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "User succesfully deleted")
}

func (u *UserController) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Here's the user selected")
}
