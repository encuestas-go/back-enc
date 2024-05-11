package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateSocioeconomicSurvey(t *testing.T) {

	t.Run("Created Socioeconomic Survey Succesfully", func(t *testing.T) {
		//GIVEN
		socioeconomic := domain.SocioeconomicStatus{
			ID:                  1,
			IDUser:              1,
			FullName:            "",
			BirthDate:           "",
			Nationality:         "",
			Gender:              "",
			Age:                 23,
			MaritalStatus:       "",
			ResidenceAddress:    "",
			ResidenceCity:       "",
			PostalCode:          62345,
			State:               "Morelos",
			SocioeconomicStatus: "",
			Language:            "",
			DegreeAspired:       "",
			LastDegreeFather:    "",
			LastDegreeMother:    "",
		}

		e := echo.New()
		socioeconomicJSON, err := json.Marshal(socioeconomic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/crear/nivelSocioeconomico", strings.NewReader(string(socioeconomicJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		socioeconomicRepository := repository.InitializeSocioeconomicRepository(db)

		mock.ExpectExec(`INSERT INTO ENCUESTA_NIVEL_SOCIOECONOMICO`).
			WithArgs(socioeconomic.IDUser, socioeconomic.FullName, socioeconomic.BirthDate,
				socioeconomic.Nationality, socioeconomic.Gender, socioeconomic.Age, socioeconomic.MaritalStatus,
				socioeconomic.ResidenceAddress, socioeconomic.ResidenceCity, socioeconomic.PostalCode,
				socioeconomic.State, socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
				socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitSocioeconomicController(socioeconomicRepository).Create((echoContext))

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusCreated,
				Message:    fmt.Sprintf("Socioeconomic Status survey succesfully created"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Create Socioeconomic survey fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/crear/nivelSocioeconomico", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitUserController(nil).Create((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestUpdateSocioeconomicSurvey(t *testing.T) {
	t.Run("Created Socioeconomic Survey Succesfully", func(t *testing.T) {
		//GIVEN
		socioeconomic := domain.SocioeconomicStatus{
			ID:                  1,
			IDUser:              1,
			FullName:            "",
			BirthDate:           "",
			Nationality:         "",
			Gender:              "",
			Age:                 23,
			MaritalStatus:       "",
			ResidenceAddress:    "",
			ResidenceCity:       "",
			PostalCode:          62345,
			State:               "Morelos",
			SocioeconomicStatus: "",
			Language:            "",
			DegreeAspired:       "",
			LastDegreeFather:    "",
			LastDegreeMother:    "",
		}

		e := echo.New()
		socioeconomicJSON, err := json.Marshal(socioeconomic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/crear/nivelSocioeconomico", strings.NewReader(string(socioeconomicJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		socioeconomicRepository := repository.InitializeSocioeconomicRepository(db)

		mock.ExpectExec(`UPDATE ENCUESTA_NIVEL_SOCIOECONOMICO`).
			WithArgs(socioeconomic.IDUser, socioeconomic.FullName, socioeconomic.BirthDate,
				socioeconomic.Nationality, socioeconomic.Gender, socioeconomic.Age, socioeconomic.MaritalStatus,
				socioeconomic.ResidenceAddress, socioeconomic.ResidenceCity, socioeconomic.PostalCode,
				socioeconomic.State, socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
				socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitSocioeconomicController(socioeconomicRepository).Create((echoContext))

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusCreated,
				Message:    fmt.Sprintf("Socioeconomic Status survey succesfully created"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

}
