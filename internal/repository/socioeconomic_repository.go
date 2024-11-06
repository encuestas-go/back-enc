package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type SocioeconomicRepositoryService struct {
	db *sql.DB
}

func InitializeSocioeconomicRepository(db *sql.DB) *SocioeconomicRepositoryService {
	return &SocioeconomicRepositoryService{
		db: db,
	}
}

func (s *SocioeconomicRepositoryService) Insert(socioeconomic domain.SocioeconomicStatus) error {
	result, err := s.db.Exec(`
	INSERT INTO ENCUESTA_NIVEL_SOCIOECONOMICO(ID_USUARIO, NOMBRE_COMPLETO, FECHA_NACIMIENTO, NACIONALIDAD, 
		SEXO, EDAD, ESTADO_CIVIL, LONGITUD, LATITUD , DIRECCION_RESIDENCIA, ESTATUS_SOCIOECONOMICO,IDIOMA, 
		GRADO_ESTUDIOS_ASPIRAR, ULTIMO_GRADO_PADRE, ULTIMO_GRADO_MADRE)
	VALUES(?,?,?,?, ?, ?, ?,?,?,?,?,?,?,?,?);
	
	`, socioeconomic.IDUser, socioeconomic.FullName, socioeconomic.BirthDate, socioeconomic.Nationality, socioeconomic.Gender,
		socioeconomic.Age, socioeconomic.MaritalStatus, socioeconomic.Longitude, socioeconomic.Latitude, socioeconomic.ResidenceAddress,
		socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
		socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_NIVEL_SOCIOECONOMICO table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_NIVEL_SOCIOECONOMICO table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_NIVEL_SOCIOECONOMICO table")
	}
	return nil
}

func (s *SocioeconomicRepositoryService) Update(socioeconomic domain.SocioeconomicStatus) error {
	result, err := s.db.Exec(`
	UPDATE ENCUESTA_NIVEL_SOCIOECONOMICO SET
                                     NOMBRE_COMPLETO = ?,
                                     FECHA_NACIMIENTO = ?,
                                     NACIONALIDAD = ?,
                                     SEXO = ?,
                                     EDAD = ?,
                                     ESTADO_CIVIL = ?,
									 LONGITUD = ?,
                                     LATITUD = ?,
                                     DIRECCION_RESIDENCIA = ?,
                                     ESTATUS_SOCIOECONOMICO = ?,
                                     IDIOMA = ?,
                                     GRADO_ESTUDIOS_ASPIRAR = ?,
                                     ULTIMO_GRADO_PADRE = ?,
                                     ULTIMO_GRADO_MADRE = ?
                                     WHERE ID_USUARIO = ?;
	`, socioeconomic.FullName, socioeconomic.BirthDate, socioeconomic.Nationality, socioeconomic.Gender,
		socioeconomic.Age, socioeconomic.MaritalStatus, socioeconomic.Longitude, socioeconomic.Latitude, socioeconomic.ResidenceAddress,
		socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
		socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother, socioeconomic.IDUser)
	if err != nil {
		log.Println("Data could not be updated into ENCUESTA_NIVEL_SOCIOECONOMICO table, the error was:", err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated into ENCUESTA_NIVEL_SOCIOECONOMICO table")
		return nil
	} else if rowsUpdated == 0 {
		return errors.New("data could not be update into ENCUESTA_NIVEL_SOCIOECONOMICO table")
	}

	return nil
}

func (s *SocioeconomicRepositoryService) Delete(idUser int) error {
	result, err := s.db.Exec(`
	DELETE FROM ENCUESTA_NIVEL_SOCIOECONOMICO WHERE ID_USUARIO = ?;
	`, idUser)
	if err != nil {
		log.Println("Could not delete the ID on ENCUESTA_NIVEL_SOCIOECONOMICO table, the error was: ", err)
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete information with the requested ID: ", err)
		return err
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from ENCUESTA_NIVEL_SOCIOECONOMICO table", idUser)
		return nil
	} else if rowsDeleted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_NIVEL_SOCIOECONOMICO table")
	}
	return nil
}

func (s *SocioeconomicRepositoryService) GetAllOrByID(userID int) ([]domain.SocioeconomicStatus, error) {
	var query = `SELECT * FROM ENCUESTA_NIVEL_SOCIOECONOMICO;`

	if userID > 0 {
		query = fmt.Sprintf(`SELECT * FROM ENCUESTA_NIVEL_SOCIOECONOMICO WHERE ID_USUARIO = %v;`, userID)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.SocioeconomicStatus{}, err
	}

	defer rows.Close()

	socioeconomicSurvey := []domain.SocioeconomicStatus{}
	for rows.Next() {
		socioeconomic := domain.SocioeconomicStatus{}
		if err = rows.Scan(&socioeconomic.ID, &socioeconomic.IDUser, &socioeconomic.FullName, &socioeconomic.BirthDate,
			&socioeconomic.Nationality, &socioeconomic.Gender, &socioeconomic.Age, &socioeconomic.MaritalStatus,
			&socioeconomic.Longitude, &socioeconomic.Latitude, &socioeconomic.ResidenceAddress, &socioeconomic.SocioeconomicStatus,
			&socioeconomic.Language, &socioeconomic.DegreeAspired, &socioeconomic.LastDegreeFather,
			&socioeconomic.LastDegreeMother); err != nil {
			return []domain.SocioeconomicStatus{}, err
		}

		socioeconomicSurvey = append(socioeconomicSurvey, socioeconomic)
	}

	return socioeconomicSurvey, nil
}
