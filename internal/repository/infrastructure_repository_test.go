package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Insert_Infrastructure(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	infrastructureRepository := InitializeInfrastructureRepository(db)

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

	mock.ExpectExec(`INSERT INTO ENCUESTA_INFRAESTRUCTURA_HOGAR`).
		WithArgs(infrastructure.UserID, infrastructure.Zone, infrastructure.Permanence, infrastructure.InfraestructureStatus,
			infrastructure.FloorType, infrastructure.RoofType, infrastructure.WallType, infrastructure.TotalMembers,
			infrastructure.TotalRooms, infrastructure.HouseholdEquipment, infrastructure.BasicServices, infrastructure.OtherProperties).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//WHEN
	err = infrastructureRepository.Insert(infrastructure)

	//THEN
	assert.NoError(t, err)
}

func Test_Update_Infrastructure(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	infrastructureRepository := InitializeInfrastructureRepository(db)

	infrastructure := domain.HouseholdInfrastructure{
		Zone:                  "Urbana",
		Permanence:            "Mediano plazo",
		InfraestructureStatus: "Regular",
		FloorType:             "Cemento",
		RoofType:              "Concreto",
		WallType:              "Concreto",
		TotalMembers:          4,
		TotalRooms:            3,
		HouseholdEquipment:    "Estufa, lavadora, television",
		BasicServices:         "Luz, Agua, Gas",
		OtherProperties:       true,
	}

	mock.ExpectExec(`UPDATE ENCUESTA_INFRAESTRUCTURA_HOGAR`).
		WithArgs(infrastructure.Zone, infrastructure.Permanence, infrastructure.InfraestructureStatus,
			infrastructure.FloorType, infrastructure.RoofType, infrastructure.WallType, infrastructure.TotalMembers,
			infrastructure.TotalRooms, infrastructure.HouseholdEquipment, infrastructure.BasicServices,
			infrastructure.OtherProperties, infrastructure.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//WHEN
	err = infrastructureRepository.Update(infrastructure)

	//THEN
	assert.NoError(t, err)
}

func Test_Delete_Infrastructure(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	infrastructureRepository := InitializeInfrastructureRepository(db)

	userID := 1

	mock.ExpectExec("DELETE FROM ENCUESTA_INFRAESTRUCTURA_HOGAR").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	//WHEN
	err = infrastructureRepository.Delete(userID)

	//THEN
	assert.NoError(t, err)
}

func Test_Get_Infrastructure(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	infrastructureRepository := InitializeInfrastructureRepository(db)

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
}
