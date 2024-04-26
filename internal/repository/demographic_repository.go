package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type DemographicRepositoryService struct {
	db *sql.DB
}

func InitializeDemographicRepository(db *sql.DB) *DemographicRepositoryService {
	return &DemographicRepositoryService{
		db: db,
	}
}

func (d *DemographicRepositoryService) Insert(demographic domain.DemographicStatus) error {
	result, err := d.db.Exec(`
	INSERT INTO ENCUESTA_NIVEL_DEMOGRAFICO(
		ID_USUARIO, TIPO_VIVIENDA, TIPO_CONDICION, TRANSPORTE_PROPIO, MONTO_INGRESOS, NUM_INTEGRANTES_TRABAJAN,
		NUM_INTEGRANTES_MENOR_EDAD, DESPENSA_MENSUAL, APOYOS_GOBIERNO)
	VALUES(?, ?, ?, ?, ?, ?, ? , ?, ?);
	`, demographic.UserID, demographic.HousingType, demographic.HouseCondition, demographic.OwnTransport, demographic.IncomeAmount,
		demographic.WorkingMembers, demographic.MembersUnderage, demographic.MonthlyExpenses, demographic.GovermentSupport)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_NIVEL_DEMOGRAFICO table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_NIVEL_DEMOGRAFICO table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_NIVEL_DEMOGRAFICO table")
	}
	return nil
}

func (d *DemographicRepositoryService) Update(demographic domain.DemographicStatus, id int) error {
	result, err := d.db.Exec(`
	UPDATE ENCUESTA_NIVEL_DEMOGRAFICO SET ID_USUARIO = ?,
                                      TIPO_VIVIENDA =  ?,
                                      TIPO_CONDICION =  ?,
                                      TRANSPORTE_PROPIO = ?,
                                      MONTO_INGRESOS = ?,
                                      NUM_INTEGRANTES_TRABAJAN = ?,
                                      NUM_INTEGRANTES_MENOR_EDAD = ?,
                                      DESPENSA_MENSUAL = ?,
                                      APOYOS_GOBIERNO = ?
                                      WHERE ID = ?;
	`, demographic.UserID, demographic.HousingType, demographic.HouseCondition, demographic.OwnTransport, demographic.IncomeAmount,
		demographic.WorkingMembers, demographic.MembersUnderage, demographic.MonthlyExpenses, demographic.GovermentSupport, id)
	if err != nil {
		log.Println("Data could not be updated into ENCUESTA_NIVEL_DEMOGRAFICO table, the error was:", err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsUpdated > 0 {
		log.Println("Data Successfully updated into ENCUESTA_NIVEL_DEMOGRAFICO table")
		return nil
	} else if rowsUpdated == 0 {
		return errors.New("data could not be updated into ENCUESTA_NIVEL_DEMOGRAFICO table")
	}
	return nil
}

func (d *DemographicRepositoryService) Delete(demographic domain.DemographicStatus, id int) error {
	result, err := d.db.Exec("DELETE FROM ENCUESTA_NIVEL_DEMOGRAFICO WHERE ID = ?;", id)
	if err != nil {
		log.Println("Could not delete the ID on ENCUESTA_NIVEL_DEMOGRAFICO table, the error was: ", err)
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete information with the requested ID: ", err)
		return err
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from ENCUESTA_NIVEL_DEMOGRAFICO table", id)
		return nil
	} else if rowsDeleted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_NIVEL_DEMOGRAFICO table")
	}
	return nil
}

func (d *DemographicRepositoryService) Get(demographic domain.DemographicStatus, id int) error {
	return nil
}
