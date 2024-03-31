package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SocioeconomicStatusController struct{}

func InitSocioeconomicController() *SocioeconomicStatusController {
	return &SocioeconomicStatusController{}
}

func (se *SocioeconomicStatusController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, " Socioeconomic Status survey succesfully created ")
}

func (se *SocioeconomicStatusController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Socioeconomic Status survey succesfully updated ")
}

func (se *SocioeconomicStatusController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Socioeconomic Status survey succesfully deleted ")
}

func (se *SocioeconomicStatusController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Socioeconomic Status survey: ")
}
