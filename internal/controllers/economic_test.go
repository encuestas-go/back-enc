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

func TestCreateEconomicSurvey(t *testing.T) {
	t.Run("Create economic survey succesfully", func(t *testing.T) {
		//GIVEN
		economic := domain.EconomicStatus{
			IDUser:                1,
			CurrentStatus:         "Empleado",
			JobTitle:              "Cajero",
			EmployerEstablishment: "Sanborns",
			EmploymentType:        "Semanal",
			Salary:                2345.00,
			AmountType:            "Semanal",
			WorkBenefitsType:      "Ninguna",
		}
		e := echo.New()
		economicJSON, err := json.Marshal(economic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/crear/nivelEconomico", strings.NewReader(string(economicJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		economicRepository := repository.InitializeEconomicRepository(db)

		mock.ExpectExec(`INSERT INTO ENCUESTA_NIVEL_ECONOMICO`).
			WithArgs(economic.IDUser, economic.CurrentStatus, economic.JobTitle, economic.EmployerEstablishment,
				economic.EmploymentType, economic.Salary, economic.AmountType, economic.WorkBenefitsType).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitEconomicController(economicRepository).Create((echoContext))

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusCreated,
				Message:    fmt.Sprintf("Economic Status survey succesfully created"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Create Economic survey fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/crear/nivelEconomico", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitEconomicController(nil).Create((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestUpdateEconomicSurvey(t *testing.T) {
	t.Run("Update economic survey succesfully", func(t *testing.T) {
		//GIVEN
		economic := domain.EconomicStatus{
			CurrentStatus:         "Empleado",
			JobTitle:              "Cajero",
			EmployerEstablishment: "Sanborns",
			EmploymentType:        "Semanal",
			Salary:                2345.00,
			AmountType:            "Semanal",
			WorkBenefitsType:      "Ninguna",
		}
		e := echo.New()
		economicJSON, err := json.Marshal(economic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/actualizar/nivelEconomico", strings.NewReader(string(economicJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		economicRepository := repository.InitializeEconomicRepository(db)

		mock.ExpectExec(`UPDATE ENCUESTA_NIVEL_ECONOMICO`).
			WithArgs(economic.CurrentStatus, economic.JobTitle, economic.EmployerEstablishment,
				economic.EmploymentType, economic.Salary, economic.AmountType, economic.WorkBenefitsType, economic.IDUser).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//THEN
		err = InitEconomicController(economicRepository).Update((echoContext))

		//WHEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusOK,
				Message:    fmt.Sprintf("Economic survey succesfully updated"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Update economic status fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPut, "/actualizar/nivelEconomico", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitEconomicController(nil).Update((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestDeleteEconomicSurvey(t *testing.T) {
	t.Run("Delete economic survey with ID succesfully", func(t *testing.T) {
		//GIVEN
		e := echo.New()

		request := httptest.NewRequest(http.MethodDelete, "/eliminar/Economico?user_id=1", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		economicRepository := repository.InitializeEconomicRepository(db)

		mock.ExpectExec(`DELETE FROM ENCUESTA_NIVEL_ECONOMICO`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		//THEN
		err = InitEconomicController(economicRepository).Delete(echoContext)

		//WHEN
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	})

	t.Run("Economic survey when ID is invalid", func(t *testing.T) {
		//GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodDelete, "/eliminar/nivelEconomico", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitEconomicController(nil).Delete((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	})
}

func TestGetEconomicSurvey(t *testing.T) {
	t.Run("Retrieve economic survey data succesfully", func(t *testing.T) {
		//GIVEN
		economic := domain.EconomicStatus{}

		economicJSON, err := json.Marshal(economic)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/consultar/nivelEconomico", strings.NewReader(string(economicJSON)))
		request.Header.Set("Content-Type", "application/json")

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		economicRepository := repository.InitializeEconomicRepository(db)

		userID := 1

		rows := sqlmock.NewRows([]string{"ID", "ID_USER", "SITUACION_ACTUAL", "NOMBRE_EMPLEO",
			"EMPRESA_ESTABLECIMIENTO", "TIPO_EMPLEO", "SALARIO", "TIPO_MONTO", "TIPO_PRESTACIONES", "FECHA"}).
			AddRow(1, 1, "Empleado", "Cajero", "Sanborns", "Semanal", 2345.00, "Semanal", "Ninguna", "2024-06-05")

		mock.ExpectQuery(`SELECT \* FROM ENCUESTA_NIVEL_ECONOMICO`).WillReturnRows(rows)

		// WHEN
		result, err := economicRepository.GetAllOrByID(userID)

		// THEN
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userID, result[0].IDUser)
	})
}
