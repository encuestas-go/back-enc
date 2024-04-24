package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type HouseInfrastructureRepositoryService struct {
	db *sql.DB
}

func InitializeInfrastructureRepository(db *sql.DB) *HouseInfrastructureRepositoryService {
	return &HouseInfrastructureRepositoryService{
		db: db,
	}
}

func (h *HouseInfrastructureRepositoryService) Insert(infrastructure domain.HouseholdInfrastructure) error {
	result, err := h.db.Exec(`
	INSERT INTO ENCUESTA_INFRAESTRUCTURA_HOGAR(
		ID_USUARIO, ZONA, PERMANENCIA, ESTADO_INFRAESTRUCTURA, TIPO_SUELO, TIPO_TECHO,TIPO_PARED, NUMERO_INTEGRANTES,
		NUMERO_HABITACIONES, EQUIPAMIENTO_HOGAR, SERVICIOS_BASICOS, OTRAS_PROPIEDADES)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, TRUE);
	`, infrastructure.UserID, infrastructure.Permanence, infrastructure.InfraestructureStatus, infrastructure.FloorType, infrastructure.RoofType,
		infrastructure.WallType, infrastructure.TotalMembers, infrastructure.TotalRooms, infrastructure.HouseholdEquipment, infrastructure.BasicServices,
		infrastructure.OtherProperties)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_INFRASTRUCTURA_HOGAR table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_INFRASTRUCTURA_HOGAR table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_INFRASTRUCTURA_HOGAR table")
	}
	return nil
}

func (h *HouseInfrastructureRepositoryService) Update(infrastructure domain.HouseholdInfrastructure, id int) error {
	result, err := h.db.Exec(`
	UPDATE ENCUESTA_INFRAESTRUCTURA_HOGAR SET ID_USUARIO = ?,
                                          ZONA = ?,
                                          PERMANENCIA = ?,
                                          ESTADO_INFRAESTRUCTURA = ?,
                                          TIPO_PARED = ?,
                                          TIPO_TECHO = ?,
                                          TIPO_SUELO = ?,
                                          NUMERO_INTEGRANTES = ?,
                                          NUMERO_HABITACIONES = ?,
                                          EQUIPAMIENTO_HOGAR = ?,
                                          SERVICIOS_BASICOS = ?,
                                          OTRAS_PROPIEDADES = ?
                                          WHERE ID = ?;
	`, infrastructure.UserID, infrastructure.Permanence, infrastructure.InfraestructureStatus, infrastructure.FloorType, infrastructure.RoofType,
		infrastructure.WallType, infrastructure.TotalMembers, infrastructure.TotalRooms, infrastructure.HouseholdEquipment, infrastructure.BasicServices,
		infrastructure.OtherProperties, id)
	if err != nil {
		log.Println("Unable to update data into ENCUESTA_INFRASTRUCTURA_HOGAR table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully updated into ENCUESTA_INFRASTRUCTURA_HOGAR table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot update data into the ENCUESTA_INFRASTRUCTURA_HOGAR table")
	}
	return nil
}

func (h *HouseInfrastructureRepositoryService) Delete(infrastructure domain.HouseholdInfrastructure, id int) error {
	result, err := h.db.Exec("DELETE FROM ENCUESTA_INFRAESTRUCTURA_HOGAR WHERE ID = ?;", id)
	if err != nil {
		log.Println("Could not delete the ID on ENCUESTA_INFRASTRUCTURA_HOGAR table, the error was: ")
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete information with the requested ID: ", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("ID %v was successfully deleted from ENCUESTA_INFRASTRUCTURA_HOGAR table", id)
		return nil
	} else if rowsInserted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_INFRASTRUCTURA_HOGAR table")
	}
	return nil
}

func (h *HouseInfrastructureRepositoryService) Get(infrastructure domain.HouseholdInfrastructure) error {
	return nil
}
