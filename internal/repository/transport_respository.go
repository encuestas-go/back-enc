package repository

import (
	"database/sql"
	"errors"
	"fmt"
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
	result, err := t.db.Exec("DELETE FROM ENCUESTA_TRANSPORTE WHERE ID_USUARIO = ?;", idUser)
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
		log.Printf("ID %v was successfully deleted from ENCUESTA_TRANSPORTE table", idUser)
		return nil
	} else if rowsDeleted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_TRANSPORTE table")
	}
	return nil
}

func (t *TransportRespositoryService) GetAllOrByID(userID int) ([]domain.TransportManagement, error) {
	var query = `SELECT * FROM ENCUESTA_TRANSPORTE;`

	if userID > 0 {
		query = fmt.Sprintf(`SELECT * FROM ENCUESTA_TRANSPORTE WHERE ID_USUARIO = %v;`, userID)
	}

	rows, err := t.db.Query(query)
	if err != nil {
		return []domain.TransportManagement{}, err
	}

	defer rows.Close()

	transportSurvey := []domain.TransportManagement{}
	for rows.Next() {
		transport := domain.TransportManagement{}
		if err = rows.Scan(&transport.ID, &transport.UserID, &transport.PrimaryTransport, &transport.SecondTransport,
			&transport.UsageFrequency, &transport.AccesiblePoints, &transport.FrequentDestination, &transport.TravelTime); err != nil {
			return []domain.TransportManagement{}, err
		}

		transportSurvey = append(transportSurvey, transport)
	}

	return transportSurvey, nil
}

func (t *TransportRespositoryService) GetMostUsedTransportReport() ([]domain.MostUsedTransportReport, error) {
	query := `
	SELECT TRANSPORTE_PRINCIPAL, COUNT(*) as CANTIDAD
	FROM ENCUESTA_TRANSPORTE
	GROUP BY TRANSPORTE_PRINCIPAL
	ORDER BY CANTIDAD ASC;
	`
	rows, err := t.db.Query(query)
	if err != nil {
		return []domain.MostUsedTransportReport{}, err
	}
	defer rows.Close()

	transportReport := []domain.MostUsedTransportReport{}
	for rows.Next() {
		report := domain.MostUsedTransportReport{}
		if err = rows.Scan(&report.PrimaryTransport, &report.Quantity); err != nil {
			return []domain.MostUsedTransportReport{}, err
		}
		transportReport = append(transportReport, report)
	}
	return transportReport, nil
}
