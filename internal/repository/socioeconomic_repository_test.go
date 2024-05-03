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
		ResidenceAddress:    "Calle Flores",
		ResidenceCity:       "Jiutepec",
		PostalCode:          67890,
		State:               "Morelos",
		SocioeconomicStatus: "Media",
		Language:            "Frances",
		DegreeAspired:       "Maestria",
		LastDegreeFather:    "Bachillerato",
		LastDegreeMother:    "Secundaria",
	}

	mock.ExpectExec(`INSERT INTO ENCUESTA_NIVEL_SOCIOECONOMICO`).
		WithArgs(socioeconomic.IDUser, socioeconomic.FullName, socioeconomic.BirthDate,
			socioeconomic.Nationality, socioeconomic.Gender, socioeconomic.Age, socioeconomic.MaritalStatus,
			socioeconomic.ResidenceAddress, socioeconomic.ResidenceCity, socioeconomic.PostalCode,
			socioeconomic.State, socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
			socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// THEN
	err = socioeconomicRepository.Insert(socioeconomic)

	// EXPECT
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
		ResidenceAddress:    "Calle Necatepec",
		ResidenceCity:       "Jiutepec",
		PostalCode:          67890,
		State:               "Morelos",
		SocioeconomicStatus: "Media",
		Language:            "Frances",
		DegreeAspired:       "Maestria",
		LastDegreeFather:    "Bachillerato",
		LastDegreeMother:    "Secundaria",
	}

	mock.ExpectExec(`UPDATE ENCUESTA_NIVEL_SOCIOECONOMICO`).
		WithArgs(socioeconomic.FullName, socioeconomic.BirthDate, socioeconomic.Nationality, socioeconomic.Gender,
			socioeconomic.Age, socioeconomic.MaritalStatus, socioeconomic.ResidenceAddress, socioeconomic.ResidenceCity,
			socioeconomic.PostalCode, socioeconomic.State, socioeconomic.SocioeconomicStatus, socioeconomic.Language,
			socioeconomic.DegreeAspired, socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother,
			socioeconomic.IDUser).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// THEN
	err = socioeconomicRepository.Update(socioeconomic)

	// EXPECT
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

	// THEN
	err = socioeconomicRepository.Delete(userID)

	// EXPECT
	assert.NoError(t, err)
}

func Test_Socioeconomic_Get(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	socioeconomicRepository := InitializeSocioeconomicRepository(db)

	userID := 1

	rows := sqlmock.NewRows([]string{"ID", "ID_USER", "FULL_NAME", "BIRTH_DATE", "NATIONALITY", "GENDER", "AGE",
		"MARITAL_STATUS", "RESIDENCE_ADDRESS", "RESIDENCE_CITY", "POSTAL_CODE", "STATE", "SOCIOECONOMIC_STATUS",
		"LANGUAGE", "DEGREE_ASPIRED", "LAST_DEGREE_FATHER", "LAST_DEGREE_MOTHER"}).
		AddRow(1, 1, "Maria Flores Flores", "25/05/2001", "Mexicana", "Femenino", 23, "Soltera", "Calle Necatepec",
			"Jiutepec", 67890, "Morelos", "Media", "Frances", "Maestria", "Bachillerato", "Secundaria")

	mock.ExpectQuery(`SELECT \* FROM ENCUESTA_NIVEL_SOCIOECONOMICO`).WillReturnRows(rows)

	// THEN
	result, err := socioeconomicRepository.GetAllOrByID(userID)

	// EXPECT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result[0].IDUser)
}
