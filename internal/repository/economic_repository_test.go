package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Economic_Insert(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	economicRepository := InitializeEconomicRepository(db)

	economic := domain.EconomicStatus{
		IDUser:                2,
		CurrentStatus:         "Empleado",
		JobTitle:              "Cajero",
		EmployerEstablishment: "Mambata",
		EmploymentType:        "Medio Tiempo",
		Salary:                1000,
		AmountType:            "Semanal",
		WorkBenefitsType:      "Ninguna",
	}

	mock.ExpectExec(`INSERT INTO ENCUESTA_NIVEL_ECONOMICO`).
		WithArgs(economic.IDUser, economic.CurrentStatus, economic.JobTitle, economic.EmployerEstablishment,
			economic.EmploymentType, economic.Salary, economic.AmountType, economic.WorkBenefitsType).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// THEN
	err = economicRepository.Insert(economic)

	// EXPECT
	assert.NoError(t, err)
}

func Test_Economic_Update(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	economicRepository := InitializeEconomicRepository(db)

	economic := domain.EconomicStatus{
		CurrentStatus:         "Empleado",
		JobTitle:              "Cajero",
		EmployerEstablishment: "Mambata",
		EmploymentType:        "Medio Tiempo",
		Salary:                1000,
		AmountType:            "Semanal",
		WorkBenefitsType:      "Ninguna",
	}

	mock.ExpectExec(`UPDATE ENCUESTA_NIVEL_ECONOMICO`).
		WithArgs(economic.CurrentStatus, economic.JobTitle, economic.EmployerEstablishment, economic.EmploymentType,
			economic.Salary, economic.AmountType, economic.WorkBenefitsType, economic.IDUser).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// THEN
	err = economicRepository.Update(economic)

	// EXPECT
	assert.NoError(t, err)

}

func Test_Economic_Delete(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	economicRepository := InitializeEconomicRepository(db)

	userID := 1

	mock.ExpectExec("DELETE FROM ENCUESTA_NIVEL_ECONOMICO").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	//WHEN
	err = economicRepository.Delete(userID)

	//THEN
	assert.NoError(t, err)
}

func Test_Economic_Get(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	economicRepository := InitializeEconomicRepository(db)

	userID := 1

	rows := sqlmock.NewRows([]string{"ID", "ID_USER", "SITUACION_ACTUAL", "NOMBRE_EMPLEO", "EMPRESA_ESTABLECIMIENTO",
		"TIPO_EMPLEO", "SALARIO", "TIPO_MONTO", "TIPO_PRESTACIONES"}).
		AddRow(1, 1, "Empleado", "Asistente Administrativo", "Rem-Z", "POr Horas", 200, "Horas", "Ninguna")

	mock.ExpectQuery(`SELECT \* FROM ENCUESTA_NIVEL_ECONOMICO`).WillReturnRows(rows)

	// WHEN
	result, err := economicRepository.GetAllOrByID(userID)

	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result[0].IDUser)
}
