package repository

import (
	"database/sql"
	"log"

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

func (f *ForumRepositoryService) InsertQuestion(question domain.Question) (int, error) {
	result, err := f.db.Exec(`
        INSERT INTO FORO_PREGUNTA(ID_USUARIO, PREGUNTA, NOMBRE)
        VALUES(?, ?, ?);
    `, question.IDUser, question.QuestionText, question.Name)
	if err != nil {
		log.Println("Unable to insert into the FORO_PREGUNTA table, the error is:", err)
		return 0, err
	}

	questionID, err := result.LastInsertId()
	if err != nil {
		log.Println("Unable to get the last inserted ID for question", err)
		return 0, err
	}

	log.Printf("Data successfully added to FORO_PREGUNTA table with ID %d", questionID)
	return int(questionID), nil
}

func (f *ForumRepositoryService) InsertAnswer(answer domain.Answer) error {
	result, err := f.db.Exec(`
        INSERT INTO FORO_RESPUESTA(ID_USUARIO, RESPUESTA, NOMBRE)
        VALUES(?, ?, ?);
    `, answer.IDUser, answer.AnswerText, answer.Name)
	if err != nil {
		log.Println("Unable to insert into the FORO_RESPUESTA table, the error is:", err)
		return err
	}

	answerID, err := result.LastInsertId()
	if err != nil {
		log.Println("Unable to get the last inserted ID for answer", err)
		return err
	}

	_, err = f.db.Exec(`
        INSERT INTO FORO_PREGUNTA_RESPUESTA(ID_PREGUNTA, ID_RESPUESTA)
        VALUES(?, ?);
    `, answer.QuestionID, answerID)
	if err != nil {
		log.Println("Unable to insert into the FORO_PREGUNTA_RESPUESTA table, the error is:", err)
		return err
	}

	log.Printf("Data successfully added to FORO_RESPUESTA and FORO_PREGUNTA_RESPUESTA tables")
	return nil
}

func (f *ForumRepositoryService) GetAll() (domain.AnswerResponseForum, error) {
	questionsQuery := `
	SELECT FORO_PREGUNTA.ID, FORO_PREGUNTA.PREGUNTA, USUARIO.NOMBRE
	FROM FORO_PREGUNTA
	INNER JOIN USUARIO
	ON FORO_PREGUNTA.ID_USUARIO = USUARIO.ID;`

	questionRows, err := f.db.Query(questionsQuery)
	if err != nil {
		return domain.AnswerResponseForum{}, err
	}
	defer questionRows.Close()

	questions := []domain.Question{}
	for questionRows.Next() {
		question := domain.Question{}
		if err = questionRows.Scan(&question.ID, &question.QuestionText, &question.Name); err != nil {
			return domain.AnswerResponseForum{}, err
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
			return domain.AnswerResponseForum{}, err
		}
		defer answerRows.Close()

		answers := []domain.Answer{}
		for answerRows.Next() {
			answer := domain.Answer{}
			if err := answerRows.Scan(&answer.ID, &answer.AnswerText, &answer.Name); err != nil {
				return domain.AnswerResponseForum{}, err
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
