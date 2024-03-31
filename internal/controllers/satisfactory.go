package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SatisfactorySurveyController struct{}

func InitSatisfactorySurveyController() *SatisfactorySurveyController {
	return &SatisfactorySurveyController{}
}

func (ss *SatisfactorySurveyController) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "Satisfactory survey succesfully created")
}

func (ss *SatisfactorySurveyController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Satisfactory survey succesfully updated")
}

func (ss *SatisfactorySurveyController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Satisfactory survey succesfully deleted")
}

func (ss *SatisfactorySurveyController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Complete information of Satisfactory survey:")
}
