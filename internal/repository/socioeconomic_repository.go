package repository

import (
	"database/sql"
	"errors"
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
	INSERT INTO ENCUESTA_NIVEL_SOCIOECONOMICO(
		ID_USUARIO, NOMBRE_COMPLETO, FECHA_NACIMIENTO,
		NACIONALIDAD, SEXO, EDAD, ESTADO_CIVIL, 
		DIRECCION_RESIDENCIA, CIUDAD_RESIDENCIA, CODIGO_POSTAL, 
		ENTIDAD_FEDERATIVA, ESTATUS_SOCIOECONOMICO, 
		IDIOMA, GRADO_ESTUDIOS_ASPIRAR, ULTIMO_GRADO_PADRE,ULTIMO_GRADO_MADRE)
    VALUES(?,?,?,?, ?,?, ?,?,?,?,?,?,?,?,?,?);
	`, socioeconomic.IDUserType, socioeconomic.FullName, socioeconomic.BirthDate, socioeconomic.Nationality, socioeconomic.Gender,
		socioeconomic.Age, socioeconomic.MaritalStatus, socioeconomic.ResidenceAddress, socioeconomic.ResidenceCity, socioeconomic.PostalCode,
		socioeconomic.State, socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
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

func (s *SocioeconomicRepositoryService) Update(socioeconomic domain.SocioeconomicStatus, id int) error {
	result, err := s.db.Exec(`
	UPDATE ENCUESTA_NIVEL_SOCIOECONOMICO SET ID_USUARIO = ?,
                                     NOMBRE_COMPLETO = ?,
                                     FECHA_NACIMIENTO = ?,
                                     NACIONALIDAD = ?,
                                     SEXO = ?,
                                     EDAD = ?,
                                     ESTADO_CIVIL = ?,
                                     DIRECCION_RESIDENCIA = ?,
                                     CIUDAD_RESIDENCIA = ?,
                                     CODIGO_POSTAL = ?,
                                     ENTIDAD_FEDERATIVA = ?,
                                     ESTATUS_SOCIOECONOMICO = ?,
                                     IDIOMA = ?,
                                     GRADO_ESTUDIOS_ASPIRAR = ?,
                                     ULTIMO_GRADO_PADRE = ?,
                                     ULTIMO_GRADO_MADRE = ?
                                     WHERE ID = ?;

	`, socioeconomic.IDUserType, socioeconomic.FullName, socioeconomic.BirthDate, socioeconomic.Nationality, socioeconomic.Gender,
		socioeconomic.Age, socioeconomic.MaritalStatus, socioeconomic.ResidenceAddress, socioeconomic.ResidenceCity, socioeconomic.PostalCode,
		socioeconomic.State, socioeconomic.SocioeconomicStatus, socioeconomic.Language, socioeconomic.DegreeAspired,
		socioeconomic.LastDegreeFather, socioeconomic.LastDegreeMother, id)
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

func (s *SocioeconomicRepositoryService) Delete(socioeconomic domain.SocioeconomicStatus, id int) error {
	result, err := s.db.Exec(`
	DELETE FROM ENCUESTA_NIVEL_SOCIOECONOMICO WHERE ID = ?;
	`, id)
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
		log.Printf("ID %v was successfully deleted from ENCUESTA_NIVEL_SOCIOECONOMICO table", id)
		return nil
	} else if rowsDeleted == 0 {
		return errors.New("could not delete the requested ID in the ENCUESTA_NIVEL_SOCIOECONOMICO table")
	}
	return nil
}

func (s *SocioeconomicRepositoryService) Get(socioeconomic domain.SocioeconomicStatus) error {
	return nil
}
