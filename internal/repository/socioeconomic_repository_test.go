package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Socioeconomic_Insert(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	socioeconomicRepository := InitializeSocioeconomicRepository(db)

	socioeconomic := domain.SocioeconomicStatus{
		IDUser:              1,
		FullName:            "Hola",
		BirthDate:           "25/05/2001",
		Nationality:         "Mexicana",
		Gender:              "Femenino",
		Age:                 23,
		MaritalStatus:       "Soltera",
		Longitude:           1.547,
		Latitude:            -2.46456,
		ResidenceAddress:    "Calle Flores",
		SocioeconomicStatus: "Media",
		Language:            "Frances",
		DegreeAspired:       "Maestria",
		LastDegreeFather:    "Bachillerato",
		LastDegreeMother:    "Secundaria",
	}

	mock.ExpectExec(`INSERT INTO ENCUESTA_NIVEL_SOCIOECONOMICO`).
		WithArgs(socioeconomic.IDUser, socioeconomic.FullName, socioeconomic.BirthDate,
			socioeconomic.Nationality, socioeconomic.Gender, socioeconomic.Age, socioeconomic.MaritalStatus,
			socioeconomic.Longitude, socioeconomic.Latitude, socioeconomic.ResidenceAddress,
			socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
			socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// WHEN
	err = socioeconomicRepository.Insert(socioeconomic)

	// THEN
	assert.NoError(t, err)
}

func Test_Socioeconomic_Update(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	socioeconomicRepository := InitializeSocioeconomicRepository(db)

	socioeconomic := domain.SocioeconomicStatus{
		FullName:            "Maria Flores Flores",
		BirthDate:           "25/05/2001",
		Nationality:         "Mexicana",
		Gender:              "Femenino",
		Age:                 23,
		MaritalStatus:       "Soltera",
		Longitude:           2.67879,
		Latitude:            5.9989464,
		ResidenceAddress:    "Calle Necatepec",
		SocioeconomicStatus: "Media",
		Language:            "Frances",
		DegreeAspired:       "Maestria",
		LastDegreeFather:    "Bachillerato",
		LastDegreeMother:    "Secundaria",
	}

	mock.ExpectExec(`UPDATE ENCUESTA_NIVEL_SOCIOECONOMICO`).
		WithArgs(socioeconomic.FullName, socioeconomic.BirthDate, socioeconomic.Nationality, socioeconomic.Gender,
			socioeconomic.Age, socioeconomic.MaritalStatus, socioeconomic.Longitude, socioeconomic.Latitude,
			socioeconomic.ResidenceAddress, socioeconomic.SocioeconomicStatus, socioeconomic.Language,
			socioeconomic.DegreeAspired, socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother,
			socioeconomic.IDUser).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// WHEN
	err = socioeconomicRepository.Update(socioeconomic)

	// THEN
	assert.NoError(t, err)
}

func Test_Socioeconomic_Delete(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	socioeconomicRepository := InitializeSocioeconomicRepository(db)

	userID := 1

	mock.ExpectExec("DELETE FROM ENCUESTA_NIVEL_SOCIOECONOMICO").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// WHEN
	err = socioeconomicRepository.Delete(userID)

	// THEN
	assert.NoError(t, err)
}

func Test_Socioeconomic_Get(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	socioeconomicRepository := InitializeSocioeconomicRepository(db)

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
}
