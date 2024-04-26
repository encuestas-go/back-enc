package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type TransportRespositoryService struct {
	db *sql.DB
}

func InitializeTransportRepository(db *sql.DB) *TransportRespositoryService {
	return &TransportRespositoryService{
		db: db,
	}
}

func (t *TransportRespositoryService) Insert(transport domain.TransportManagement) error {
	result, err := t.db.Exec(`
	INSERT INTO ENCUESTA_TRANSPORTE(
		ID_USUARIO, TRANSPORTE_PRINCIPAL, TRANSPORTE_SECUNDARIO, FRECUENCIA_USO, PUNTOS_ACCESIBLES, 
		LUGAR_DESTINO_FRECUENTE, TIEMPO_TRASLADO)
    VALUES(?, ?, ?, ?, ?, ?, ?);
	`, transport.UserID, transport.PrimaryTransport, transport.SecondTransport, transport.UsageFrequency,
		transport.AccesiblePoints, transport.FrequentDestination, transport.TravelTime)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_TRANSPORTE table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_TRANSPORTE table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_TRANSPORTE table")
	}
	return nil
}

func (t *TransportRespositoryService) Update(transport domain.TransportManagement) error {
	result, err := t.db.Exec(`
	UPDATE ENCUESTA_TRANSPORTE SET 
                               TRANSPORTE_PRINCIPAL = ?,
                               TRANSPORTE_SECUNDARIO = ?,
                               FRECUENCIA_USO = ?,
                               PUNTOS_ACCESIBLES = ?,
                               LUGAR_DESTINO_FRECUENTE = ?,
                               TIEMPO_TRASLADO = ?
                               WHERE ID_USUARIO = ?;
	`, transport.PrimaryTransport, transport.SecondTransport, transport.UsageFrequency,
		transport.AccesiblePoints, transport.FrequentDestination, transport.TravelTime, transport.UserID)
	if err != nil {
		log.Println("Data could not be updated into ENCUESTA_TRANSPORTE table, the error was:", err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated into ENCUESTA_TRANSPORTE table")
		return nil
	} else if rowsUpdated == 0 {
		return errors.New("data could not be update into ENCUESTA_TRANSPORTE table")
	}
	return nil
}

func (t *TransportRespositoryService) Delete(idUser int) error {
	result, err := t.db.Exec("DELETE FROM ENCUESTA_TRANSPORTE WHERE ID_USUARIO =?;")
	if err != nil {
		log.Println("Could not delete the ID on ENCUESTA_TRANSPORTE table, the error was: ", err)
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete information with the requested ID: ", err)
		return err
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from ENCUESTA_NIVEL_ECONOMICO table", idUser)
		return nil
	} else if rowsDeleted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_NIVEL_ECONOMICO table")
	}
	return nil
}

func (t *TransportRespositoryService) Get(domain.TransportManagement) error {
	return nil
}
