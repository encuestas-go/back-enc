package controllers

import "github.com/labstack/echo/v4"

// UserController its a struct to
type UserController struct{}

func InitUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(c echo.Context) error {

	return c.JSON(200, "")
}
