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

func Test_Create_Infraestructure(t *testing.T) {
	t.Run("Create House Infraestructure survey succesfully", func(t *testing.T) {
		//GIVEN
		infrastructure := domain.HouseholdInfrastructure{
			UserID:                1,
			Zone:                  "Urbana",
			Permanence:            "Largo plazo",
			InfraestructureStatus: "Regular",
			FloorType:             "Cemento",
			RoofType:              "Concreto",
			WallType:              "Concreto",
			TotalMembers:          4,
			TotalRooms:            3,
			HouseholdEquipment:    "Estufa, lavadora, television",
			BasicServices:         "Luz, Agua, Gas",
			OtherProperties:       false,
		}

		e := echo.New()
		infrastructureJSON, err := json.Marshal(infrastructure)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/crear/InfraestructuraCasa", strings.NewReader(string(infrastructureJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		infrastructureRepository := repository.InitializeInfrastructureRepository(db)

		mock.ExpectExec(`INSERT INTO ENCUESTA_INFRAESTRUCTURA_HOGAR`).
			WithArgs(infrastructure.UserID, infrastructure.Zone, infrastructure.Permanence, infrastructure.InfraestructureStatus,
				infrastructure.FloorType, infrastructure.RoofType, infrastructure.WallType, infrastructure.TotalMembers,
				infrastructure.TotalRooms, infrastructure.HouseholdEquipment, infrastructure.BasicServices, infrastructure.OtherProperties).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitHouseInfrastructureController(infrastructureRepository).Create(echoContext)

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusCreated,
				Message:    fmt.Sprintf("Household Infrastructure survey succesfully created"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Create House Infraestructure survey fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/crear/InfraestructuraCasa", strings.NewReader(`{"}`))
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

func Test_Update_Infraestructure(t *testing.T) {
	t.Run("Update House Infraestructure survey succesfully", func(t *testing.T) {
		//GIVEN
		infrastructure := domain.HouseholdInfrastructure{
			Zone:                  "Urbana",
			Permanence:            "Largo plazo",
			InfraestructureStatus: "Regular",
			FloorType:             "Cemento",
			RoofType:              "Concreto",
			WallType:              "Concreto",
			TotalMembers:          4,
			TotalRooms:            3,
			HouseholdEquipment:    "Estufa, lavadora, television",
			BasicServices:         "Luz, Agua, Gas",
			OtherProperties:       false,
		}

		e := echo.New()
		infrastructureJSON, err := json.Marshal(infrastructure)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPut, "/actualizar/InfraestructuraCasa", strings.NewReader(string(infrastructureJSON)))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		infrastructureRepository := repository.InitializeInfrastructureRepository(db)

		mock.ExpectExec(`UPDATE ENCUESTA_INFRAESTRUCTURA_HOGAR`).
			WithArgs(infrastructure.Zone, infrastructure.Permanence, infrastructure.InfraestructureStatus,
				infrastructure.FloorType, infrastructure.RoofType, infrastructure.WallType, infrastructure.TotalMembers,
				infrastructure.TotalRooms, infrastructure.HouseholdEquipment, infrastructure.BasicServices,
				infrastructure.OtherProperties, infrastructure.UserID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		//WHEN
		err = InitHouseInfrastructureController(infrastructureRepository).Update(echoContext)

		//THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusOK,
				Message:    fmt.Sprintf("Household Infrastructure survey succesfully updated"),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})
	t.Run("Update House Infraestructure survey fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPut, "/actualizar/InfraestructuraCasa", strings.NewReader(`{"}`))
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

func Test_Delete_Infraestructure(t *testing.T) {
	t.Run("Delete with IDUser House Infrastructure Survey Succesfully", func(t *testing.T) {
		//GIVEN
		e := echo.New()

		request := httptest.NewRequest(http.MethodDelete, "/eliminar/InfraestructuraCasa?user_id=1", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		infrastructureRepository := repository.InitializeInfrastructureRepository(db)

		mock.ExpectExec(`DELETE FROM ENCUESTA_INFRAESTRUCTURA_HOGAR`).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		//THEN
		err = InitHouseInfrastructureController(infrastructureRepository).Delete(echoContext)

		//WHEN
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	})
	t.Run("Delete House Infrastructure Survey when ID is invalid", func(t *testing.T) {
		//GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodDelete, "/eliminar/InfraestructuraCasa", nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitHouseInfrastructureController(nil).Delete((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	})
}

func Test_Get_HouseInfrastructure(t *testing.T) {
	t.Run("Retrieve house infrastructure survey data succesfully", func(t *testing.T) {
		//GIVEN
		infrastructure := domain.HouseholdInfrastructure{}

		infrastructureJSON, err := json.Marshal(infrastructure)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodGet, "/consultar/InfraestructuraCasa", strings.NewReader(string(infrastructureJSON)))
		request.Header.Set("Content-Type", "application/json")

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		infrastructureRepository := repository.InitializeInfrastructureRepository(db)

		userID := 1

		rows := sqlmock.NewRows([]string{"ID", "ID_USER", "ZONA", "PERMANENCIA", "ESTADO_INFRAESTRUCTURA",
			"TIPO_SUELO", "TIPO_TECHO", "TIPO_PARED", "NUMERO_INTEGRANTES", "NUMERO_HABITACIONES",
			"EQUIPAMIENTO_HOGAR", "SERVICIOS_BASICOS", "OTRAS_PROPIEDADES"}).
			AddRow(1, 1, "Urbana", "Mediano plazo", "Regular", "Cemento", "Concreto", "Concreto", 4, 3,
				"Estufa, lavadora, television", "Luz, Agua, Gas", true)

		mock.ExpectQuery(`SELECT \* FROM ENCUESTA_INFRAESTRUCTURA_HOGAR`).WillReturnRows(rows)

		// WHEN
		result, err := infrastructureRepository.GetAllOrByID(userID)

		// THEN
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userID, result[0].UserID)
	})
}
