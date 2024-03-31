package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SocioeconomicController struct{}

func InitSocioeconomicController() *SocioeconomicController {
	return &SocioeconomicController{}
}

func (s *SocioeconomicController) CreateSocioeconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, " Socioeconomic Status survey succesfully created ")
}

func (s *SocioeconomicController) UpdateSocioeconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Socioeconomic Status survey succesfully updated ")
}

func (s *SocioeconomicController) DeleteSocioeconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Socioeconomic Status survey succesfully deleted ")
}

func (s *SocioeconomicController) GetSocioeconomicStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Socioeconomic Status survey: ")
}
