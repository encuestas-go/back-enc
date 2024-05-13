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

func TestCreateTransportSurvey(t *testing.T) {
	t.Run("Update Transport Survey succesfully", func(t *testing.T) {
		//GIVEN
		transport := domain.TransportManagement{
			UserID:              1,
			PrimaryTransport:    "Ruta",
			SecondTransport:     "Taxi",
			UsageFrequency:      "Diario",
			AccesiblePoints:     true,
			FrequentDestination: "Escuela",
			TravelTime:          "Una hora",
		}
		e := echo.New()
		transportJSON, err := json.Marshal(transport)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/crear/medioTransporte", strings.NewReader(string(transportJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		transportRepository := repository.InitializeTransportRepository(db)

		mock.ExpectExec(`INSERT INTO ENCUESTA_TRANSPORTE`).
			WithArgs(transport.UserID, transport.PrimaryTransport, transport.SecondTransport, transport.UsageFrequency,
				transport.AccesiblePoints, transport.FrequentDestination, transport.TravelTime).
			WillReturnResult(sqlmock.NewResult(1, 1))
		//WHEN
		err = InitTransportController(transportRepository).Create(echoContext)

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusCreated,
				Message:    fmt.Sprintf("Transport Management survey succesfully created"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Create Transport survey fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/crear/medioTransporte", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitTransportController(nil).Create((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestUpdateTransportSurvey(t *testing.T) {
	t.Run("Update Transport Survey succesfully", func(t *testing.T) {
		//GIVEN
		//GIVEN
		transport := domain.TransportManagement{
			PrimaryTransport:    "Ruta",
			SecondTransport:     "Taxi",
			UsageFrequency:      "Diario",
			AccesiblePoints:     true,
			FrequentDestination: "Escuela",
			TravelTime:          "Una hora",
		}
		e := echo.New()
		transportJSON, err := json.Marshal(transport)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/actualizar/medioTransporte", strings.NewReader(string(transportJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		transportRepository := repository.InitializeTransportRepository(db)

		mock.ExpectExec(`UPDATE ENCUESTA_TRANSPORTE`).
			WithArgs(transport.PrimaryTransport, transport.SecondTransport, transport.UsageFrequency,
				transport.AccesiblePoints, transport.FrequentDestination, transport.TravelTime, transport.UserID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		//WHEN
		err = InitTransportController(transportRepository).Update(echoContext)

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusOK,
				Message:    fmt.Sprintf("Transport Management survey succesfully updated"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Update Transport Management,  fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPut, "/actualizar/medioTransporte", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitTransportController(nil).Update((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestDeleteTransportSurvey(t *testing.T) {
	t.Run("Delete Transport Survey with ID succesfully", func(t *testing.T) {
		//GIVEN
		e := echo.New()

		request := httptest.NewRequest(http.MethodDelete, "/eliminar/medioTransporte?user_id=1", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		transportRepository := repository.InitializeTransportRepository(db)

		mock.ExpectExec(`DELETE FROM ENCUESTA_TRANSPORTE`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		//WHEN
		err = InitTransportController(transportRepository).Delete(echoContext)

		//THEN
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	})

	t.Run("Delete Transport Survey with invalid ID", func(t *testing.T) {
		//GIVEN
		e := echo.New()

		request := httptest.NewRequest(http.MethodDelete, "/eliminar/medioTransporte", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		//WHEN
		err := InitTransportController(nil).Delete((echoContext))

		//THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)

	})
}

func TestGetTransportSurvey(t *testing.T) {
	t.Run("Retrieve information of Transport Survey succesfully", func(t *testing.T) {
		//GIVEN
		transport := domain.TransportManagement{}
		//e := echo.New()
		transportJSON, err := json.Marshal(transport)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/consultar/medioTransporte", strings.NewReader(string(transportJSON)))
		request.Header.Set("Content-Type", "application/json")

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		transportRepository := repository.InitializeTransportRepository(db)

		rows := sqlmock.NewRows([]string{"ID", "ID_USER", "TRANSPORTE_PRINCIPAL", "TRANSPORTE_SECUNDARIO",
			"FRECUENCIA_USO", "PUNTOS_ACCESIBLES", "LUGAR_DESTINO_FRECUENTE", "TIEMPO_TRASLADO"}).
			AddRow(1, 1, "Ruta", "Taxi", "Diario", true, "Escuela", "Una hora")

		mock.ExpectQuery(`SELECT \* FROM ENCUESTA_TRANSPORTE`).WillReturnRows(rows)

		userID := 1
		//WHEN
		result, err := transportRepository.GetAllOrByID(userID)
		//THEN
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userID, result[0].UserID)
	})
}
