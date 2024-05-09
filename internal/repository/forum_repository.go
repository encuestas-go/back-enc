package repository

import (
	"database/sql"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type ForumRepositoryService struct {
	db *sql.DB
}

func InitializeForumRepository(db *sql.DB) *ForumRepositoryService {
	return &ForumRepositoryService{
		db: db,
	}
}

func (f *ForumRepositoryService) GetAll() ([]domain.AnswerResponseForum, error) {
	var query = `SELECT FORO_PREGUNTA.ID, FORO_PREGUNTA.PREGUNTA, USUARIO.NOMBRE
	FROM FORO_PREGUNTA
	INNER JOIN USUARIO
	ON FORO_PREGUNTA.ID_USUARIO = USUARIO.ID;`

	rows, err := f.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	questions := []domain.Question{}
	for rows.Next() {
		question := domain.Question{}
		if err = rows.Scan(&question.ID, &question.IDUser, &question.QuestionText, &question.Name); err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}
