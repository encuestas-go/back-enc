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

func (f *ForumRepositoryService) GetAll() (domain.AnswerResponseForum{}, error) {
	questionsQuery := `
	SELECT FORO_PREGUNTA.ID, FORO_PREGUNTA.PREGUNTA, USUARIO.NOMBRE
	FROM FORO_PREGUNTA
	INNER JOIN USUARIO
	ON FORO_PREGUNTA.ID_USUARIO = USUARIO.ID;`

	questionRows, err := f.db.Query(questionsQuery)
	if err != nil {
		return nil, err
	}
	defer questionRows.Close()

	questions := []domain.Question{}
	for questionRows.Next() {
		question := domain.Question{}
		if err = questionRows.Scan(&question.ID, &question.IDUser, &question.QuestionText, &question.Name); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	//For range to search all the answers corresponding to the ID question of the associate question
	for i, question := range questions {
		answersQuery := `
		SELECT FORO_PREGUNTA.ID ,FORO_RESPUESTA.RESPUESTA, USUARIO.NOMBRE AS USUARIO_RESPUESTA
		FROM FORO_PREGUNTA_RESPUESTA
		INNER JOIN FORO_PREGUNTA  ON FORO_PREGUNTA_RESPUESTA.ID_PREGUNTA = FORO_PREGUNTA.ID
		INNER JOIN FORO_RESPUESTA  ON FORO_PREGUNTA_RESPUESTA.ID_RESPUESTA = FORO_RESPUESTA.ID
		INNER JOIN USUARIO ON FORO_RESPUESTA.ID_USUARIO = USUARIO.ID
		WHERE FORO_PREGUNTA.ID = ?;`

		answerRows, err := f.db.Query(answersQuery, question.ID)
		if err != nil {
			return nil, err
		}
		defer answerRows.Close()

		answers := []domain.Answer{}
		for answerRows.Next() {
			answer := domain.Answer{}
			if err := answerRows.Scan(&answer.ID, &answer.AnswerText, &answer.Name); err != nil {
				return nil, err
			}
			answers = append(answers, answer)
		}
		questions[i].Answers = answers
	}

	//Include the questions and answers to []AnswerResponseForum
	forumResponse := domain.AnswerResponseForum{
		Questions: questions,
	}

	return forumResponse, nil
}
