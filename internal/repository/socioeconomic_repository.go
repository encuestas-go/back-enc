package repository

import (
	"database/sql"
)

type SocioeconomicRepositoryService struct {
	db *sql.DB
}

func InitializeSocioeconomicRepository(db *sql.DB) *SocioeconomicRepositoryService {
	return &SocioeconomicRepositoryService{
		db: db,
	}
}

func (s *SocioeconomicRepositoryService) Insert() error {
	return nil
}
