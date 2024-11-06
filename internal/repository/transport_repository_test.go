package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Insert_Transport(t *testing.T) {
	//GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	transportRepository := InitializeTransportRepository(db)
	transport := domain.TransportManagement{
		UserID:              1,
		PrimaryTransport:    "Caminando",
		SecondTransport:     "Ruta",
		UsageFrequency:      "Dos veces por semana",
		AccesiblePoints:     false,
		FrequentDestination: "Universidad",
		TravelTime:          "Menos de 1 hora",
	}

	mock.ExpectExec(`INSERT INTO ENCUESTA_TRANSPORTE`).
		WithArgs(transport.UserID, transport.PrimaryTransport, transport.SecondTransport, transport.UsageFrequency,
			transport.AccesiblePoints, transport.FrequentDestination, transport.TravelTime).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//WHEN
	err = transportRepository.Insert(transport)

	//THEN
	assert.NoError(t, err)
}

func Test_Update_Transport(t *testing.T) {
	//GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	transportRepository := InitializeTransportRepository(db)
	transport := domain.TransportManagement{
		PrimaryTransport:    "Caminando",
		SecondTransport:     "Ruta",
		UsageFrequency:      "Dos veces por semana",
		AccesiblePoints:     true,
		FrequentDestination: "Universidad",
		TravelTime:          "Menos de 1 hora",
	}

	mock.ExpectExec(`UPDATE ENCUESTA_TRANSPORTE`).
		WithArgs(transport.PrimaryTransport, transport.SecondTransport, transport.UsageFrequency,
			transport.AccesiblePoints, transport.FrequentDestination, transport.TravelTime, transport.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	//WHEN
	err = transportRepository.Update(transport)

	//THEN
	assert.NoError(t, err)
}

func Test_Delete_Transport(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	transportRepository := InitializeTransportRepository(db)

	IDUser := 2

	mock.ExpectExec("DELETE FROM ENCUESTA_TRANSPORTE WHERE ID_USUARIO = ?").
		WithArgs(IDUser).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// WHEN
	err = transportRepository.Delete(IDUser)

	// THEN
	assert.NoError(t, err)
}

func Test_Get_Transport(t *testing.T) {
	//GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	transportRepository := InitializeTransportRepository(db)

	userID := 1
	rows := sqlmock.NewRows([]string{"ID", "ID_USER", "TRANSPORTE_PRINCIPAL", "TRANSPORTE_SECUNDARIO",
		"FRECUENCIA_USO", "PUNTOS_ACCESIBLES", "LUGAR_DESTINO_FRECUENTE", "TIEMPO_TRASLADO", "FECHA"}).
		AddRow(1, 1, "Caminando", "Ruta", "Dos veces por semana", true, "Universidad", "Menos de 1 hora", "2024-06-05")

	mock.ExpectQuery(`SELECT \* FROM ENCUESTA_TRANSPORTE`).WillReturnRows(rows)
	//WHEN
	result, err := transportRepository.GetAllOrByID(userID)

	//THEN
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result[0].UserID)
}
