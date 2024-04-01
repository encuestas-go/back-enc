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

func (u *UserController) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Succesfully login  ")
}

func (u *UserController) LogOut(c echo.Context) error {
	return c.JSON(http.StatusOK, "Succesfully log out")
}

func (u *UserController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "User succesfully created ")
}

func (u *UserController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "User succesfully updated")
}

func (u *UserController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "User succesfully deleted")
}

func (u *UserController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Here's the user selected")
}
