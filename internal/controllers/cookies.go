package controllers

import "github.com/labstack/echo/v4"

func IsAdminUser(c echo.Context) bool {
	c.C

	return true
}
