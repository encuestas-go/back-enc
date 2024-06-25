package repository

import (
	"database/sql"
	"errors"
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

func (s *SatisfactorySurveyRepositoryService) Get() (domain.SatisfactorySurvey, error) {
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
		if err == sql.ErrNoRows {
			return survey, nil
		}
		return survey, err
	}

	return survey, nil
}
