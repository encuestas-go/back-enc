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
	forum := []domain.AnswerResponseForum{}

	return forum, nil
}
