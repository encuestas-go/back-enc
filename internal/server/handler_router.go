package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const ADMINISTRATOR = 1

func NeedAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookieIDTypeUser, err := c.Cookie("id_type_user")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
				Message: "Unauthorized user, need to be admin",
			})
		}

		IDTypeUserConverted, err := strconv.Atoi(cookieIDTypeUser.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
				Message: "Unauthorized user, need to be admin",
			})
		}

		if IDTypeUserConverted == 0 {
			return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
				Message: "Unauthorized user, need to be admin",
			})
		}

		if IDTypeUserConverted != ADMINISTRATOR {
			return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
				Message: "Unauthorized user, need to be admin",
			})
		}

		return next(c)
	}
}
