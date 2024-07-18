package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type SatisfactorySurveyRepositoryService struct {
	db *sql.DB
}

func InitializeSatisfactorySurveyRepository(db *sql.DB) *SatisfactorySurveyRepositoryService {
	return &SatisfactorySurveyRepositoryService{
		db: db,
	}
}

func (s *SatisfactorySurveyRepositoryService) InsertLikert(survey domain.SatisfactoryLikertSurvey) error {
	result, err := s.db.Exec(`
	INSERT INTO 
	    ENCUESTA_SATISFACCION_LINKERT(FACILIDAD_USO, CLARIDAD_INSTRUCCION, RELEVANCIA_CONTENIDO, 
	                                  RAPIDEZ_CARGA, SATISFACCCION_GENERAL, ID_USUARIO) 
	    VALUE (?,?,?,?,?,?);`, survey.FacilidadUso, survey.ClaridadInstruccion, survey.RelevanciaContenido,
		survey.RapidezCarga, survey.SatisfaccionGeneral, survey.IDUsuario)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_SATISFACCION table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtain values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_SATISFACCION table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_SATISFACCION table")
	}
	return nil
}

func (s *SatisfactorySurveyRepositoryService) Insert(survey domain.SatisfactorySurvey) error {
	result, err := s.db.Exec(`
		INSERT INTO ENCUESTA_SATISFACCION (ID_USUARIO, FECHA_PROGRAMADA)
		VALUES (?, ?);
	`, survey.IDUser, survey.ScheduledDate)
	if err != nil {
		log.Println("Unable to insert into the ENCUESTA_SATISFACCION table, the error is:", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtain values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to ENCUESTA_SATISFACCION table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the ENCUESTA_SATISFACCION table")
	}
	return nil
}

func (s *SatisfactorySurveyRepositoryService) Get(userID int) ([]domain.SatisfactoryLikertSurvey, error) {
	var query = ` SELECT * FROM ENCUESTA_SATISFACCION_LINKERT;`

	if userID > 0 {
		query = fmt.Sprintf(`SELECT * FROM ENCUESTA_SATISFACCION_LINKERT WHERE ID_USUARIO = %v;`, userID)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.SatisfactoryLikertSurvey{}, err
	}

	surveys := []domain.SatisfactoryLikertSurvey{}
	for rows.Next() {
		var survey domain.SatisfactoryLikertSurvey
		if err = rows.Scan(&survey.ID, &survey.FacilidadUso, &survey.ClaridadInstruccion, &survey.RelevanciaContenido,
			&survey.RapidezCarga, &survey.SatisfaccionGeneral, &survey.IDUsuario); err != nil {
			return nil, err
		}
		surveys = append(surveys, survey)
	}

	return surveys, nil
}

func (s SatisfactorySurveyRepositoryService) Update(survey domain.SatisfactoryLikertSurvey) error {
	result, err := s.db.Exec(`
		UPDATE ENCUESTA_SATISFACCION_LINKERT SET 
                                         FACILIDAD_USO = ?,
                                         CLARIDAD_INSTRUCCION = ?,
                                         RELEVANCIA_CONTENIDO = ?,
                                         RAPIDEZ_CARGA = ?,
                                         SATISFACCCION_GENERAL = ? WHERE ID_USUARIO = ?;
	`, &survey.FacilidadUso, &survey.ClaridadInstruccion, &survey.RelevanciaContenido, &survey.RapidezCarga,
		&survey.SatisfaccionGeneral, &survey.IDUsuario)
	if err != nil {
		log.Println("Data could not be updated into ENCUESTA_SATISFACCION table, the error was:", err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated into ENCUESTA_SATISFACCION table")
		return nil
	} else if rowsUpdated == 0 {
		return errors.New("data could not be update into ENCUESTA_SATISFACCION table")
	}
	return nil
}

func (s SatisfactorySurveyRepositoryService) Delete(userID int) error {
	res, err := s.db.Exec(`DELETE FROM ENCUESTA_SATISFACCION_LINKERT WHERE ID_USUARIO = ?;`, userID)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("cannot delete a cultural activity")
	}

	return nil
}

func (s *SatisfactorySurveyRepositoryService) GetSchedule() (domain.SatisfactorySurvey, error) {
	query := `
	SELECT ID, ID_USUARIO, FECHA_PROGRAMADA
	FROM ENCUESTA_SATISFACCION
	ORDER BY ID DESC
	LIMIT 1;
	`
	row := s.db.QueryRow(query)

	survey := domain.SatisfactorySurvey{}
	err := row.Scan(&survey.ID, &survey.IDUser, &survey.ScheduledDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return survey, nil
		}
		return survey, err
	}

	return survey, nil
}
