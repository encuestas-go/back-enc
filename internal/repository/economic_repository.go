package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type EconomicRepositoryService struct {
	db *sql.DB
}

func InitializeEconomicRepository(db *sql.DB) *EconomicRepositoryService {
	return &EconomicRepositoryService{
		db: db,
	}
}

func (e *EconomicRepositoryService) Insert(economic domain.EconomicStatus) error {
	result, err := e.db.Exec(`
	INSERT INTO ENCUESTA_NIVEL_ECONOMICO(
		ID_USUARIO, SITUACION_ACTUAL, NOMBRE_EMPLEO, EMPRESA_ESTABLECIMIENTO, TIPO_EMPLEO, 
		SALARIO, TIPO_MONTO, TIPO_PRESTACIONES)
    VALUES(?, ?, ?, ?, ?, ?, ?, ?);
	`, economic.IDUser, economic.CurrentStatus, economic.JobTitle, economic.EmployerEstablishment,
		economic.EmploymentType, economic.Salary, economic.AmountType, economic.WorkBenefitsType)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_NIVEL_ECONOMICO table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_NIVEL_ECONOMICO table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_NIVEL_ECONOMICO table")
	}
	return nil
}

func (e *EconomicRepositoryService) Update(economic domain.EconomicStatus, id int) error {
	result, err := e.db.Exec(`
	UPDATE ENCUESTA_NIVEL_ECONOMICO SET ID_USUARIO = ?,
                                    SITUACION_ACTUAL = ?,
                                    NOMBRE_EMPLEO = ?,
                                    EMPRESA_ESTABLECIMIENTO = ?,
                                    TIPO_EMPLEO = ?,
                                    SALARIO = ?,
                                    TIPO_MONTO = ?,
                                    TIPO_PRESTACIONES = ?
                                    WHERE ID = ?;
	`, economic.IDUser, economic.CurrentStatus, economic.JobTitle, economic.EmployerEstablishment,
		economic.EmploymentType, economic.Salary, economic.AmountType, economic.WorkBenefitsType, id)
	if err != nil {
		log.Println("Data could not be updated into ENCUESTA_NIVEL_ECONOMICO table, the error was:", err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated into ENCUESTA_NIVEL_ECONOMICO table")
		return nil
	} else if rowsUpdated == 0 {
		return errors.New("data could not be update into ENCUESTA_NIVEL_ECONOMICO table")
	}
	return nil
}

func (e *EconomicRepositoryService) Delete(economic domain.EconomicStatus, id int) error {
	result, err := e.db.Exec("DELETE FROM ENCUESTA_NIVEL_ECONOMICO WHERE ID =?;", id)
	if err != nil {
		log.Println("Could not delete the ID on ENCUESTA_NIVEL_ECONOMICO table, the error was: ", err)
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete information with the requested ID: ", err)
		return err
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from ENCUESTA_NIVEL_ECONOMICO table", id)
		return nil
	} else if rowsDeleted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_NIVEL_ECONOMICO table")
	}
	return nil
}

func (e *EconomicRepositoryService) Get(economic domain.EconomicStatus) error {
	return nil
}
