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
			IDUser:              1,
			FullName:            "Paula Hierro Narvarez",
			BirthDate:           "30/04/2000",
			Nationality:         "Mexicana",
			Gender:              "Femenino",
			Age:                 23,
			MaritalStatus:       "Soltera",
			Longitude:           2.445,
			Latitude:            2.55989,
			ResidenceAddress:    "Calle Juarez",
			SocioeconomicStatus: "Media-Baja",
			Language:            "Ruso",
			DegreeAspired:       "Doctorado",
			LastDegreeFather:    "Secundaria",
			LastDegreeMother:    "Primaria",
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
				socioeconomic.Longitude, socioeconomic.Latitude, socioeconomic.ResidenceAddress,
				socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
				socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitSocioeconomicController(socioeconomicRepository).Create(echoContext)

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
	t.Run("Updated Socioeconomic Survey Succesfully", func(t *testing.T) {
		//GIVEN
		socioeconomic := domain.SocioeconomicStatus{
			ID:                  1,
			FullName:            "Paula Hierro Narvarez",
			BirthDate:           "30/04/2000",
			Nationality:         "Mexicana",
			Gender:              "Femenino",
			Age:                 24,
			MaritalStatus:       "Soltera",
			Longitude:           2.43564,
			Latitude:            2.4456,
			ResidenceAddress:    "Calle Juarez",
			SocioeconomicStatus: "Media-Baja",
			Language:            "Ruso",
			DegreeAspired:       "Doctorado",
			LastDegreeFather:    "Secundaria",
			LastDegreeMother:    "Secundaria",
		}

		e := echo.New()
		socioeconomicJSON, err := json.Marshal(socioeconomic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/actualizar/nivelSocioeconomico", strings.NewReader(string(socioeconomicJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		socioeconomicRepository := repository.InitializeSocioeconomicRepository(db)

		mock.ExpectExec(`UPDATE ENCUESTA_NIVEL_SOCIOECONOMICO`).
			WithArgs(socioeconomic.FullName, socioeconomic.BirthDate,
				socioeconomic.Nationality, socioeconomic.Gender, socioeconomic.Age, socioeconomic.MaritalStatus,
				socioeconomic.Longitude, socioeconomic.Latitude, socioeconomic.ResidenceAddress,
				socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
				socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother, socioeconomic.IDUser).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitSocioeconomicController(socioeconomicRepository).Update((echoContext))

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusOK,
				Message:    fmt.Sprintf("Socioeconomic survey succesfully updated"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Update socioeconomic status fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPut, "/actualizar/nivelSocioeconomico", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitSocioeconomicController(nil).Update((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestDeleteSocioeconomicSurvey(t *testing.T) {
	t.Run("Delete Socioeconomic Survey Succesfully", func(t *testing.T) {
		//GIVEN
		e := echo.New()

		request := httptest.NewRequest(http.MethodDelete, "/eliminar/nivelSocioeconomico?user_id=1", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		socioeconomicRepository := repository.InitializeSocioeconomicRepository(db)

		mock.ExpectExec(`DELETE FROM ENCUESTA_NIVEL_SOCIOECONOMICO`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		//WHEN
		err = InitSocioeconomicController(socioeconomicRepository).Delete(echoContext)

		//THEN
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	})

	t.Run("Delete Socioeconomic survey when ID is invalid", func(t *testing.T) {
		//GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodDelete, "/eliminar/nivelSocioeconomico", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitSocioeconomicController(nil).Delete((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	})
}

func TestGetSocioeconomicSurvey(t *testing.T) {
	t.Run("Retrieve Socioeconomic Survey data Successfully", func(t *testing.T) {
		//GIVEN
		socioeconomic := domain.SocioeconomicStatus{}
		//e := echo.New()
		socioeconomicJSON, err := json.Marshal(socioeconomic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/consultar/nivelSocioeconomico", strings.NewReader(string(socioeconomicJSON)))
		request.Header.Set("Content-Type", "application/json")

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		socioeconomicRepository := repository.InitializeSocioeconomicRepository(db)

		userID := 1

		rows := sqlmock.NewRows([]string{"ID", "ID_USER", "NOMBRE_COMPLETO", "FECHA_NACIMIENTO", "NACIONALIDAD",
			"SEXO", "EDAD", "ESTADO_CIVIL", "LONGITUD", "LATITUD", "DIRECCION_RESIDENCIA", "ESTATUS_SOCIOECONOMICO",
			"IDIOMA", "GRADO_ESTUDIOS_ASPIRAR", "ULTIMO_GRADO_PADRE", "ULTIMO_GRADO_MADRE"}).
			AddRow(1, 1, "Maria Flores Flores", "25/05/2001", "Mexicana", "Femenino", 23, "Soltera", 1.567, 3.658, "Calle Necatepec",
				"Media", "Frances", "Maestria", "Bachillerato", "Secundaria")

		mock.ExpectQuery(`SELECT \* FROM ENCUESTA_NIVEL_SOCIOECONOMICO`).WillReturnRows(rows)

		// WHEN
		result, err := socioeconomicRepository.GetAllOrByID(userID)

		// THEN
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userID, result[0].IDUser)
	})
}
