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

func (s *SocioeconomicRepositoryService) Update() error {
	return nil
}

func (s *SocioeconomicRepositoryService) Delete() error {
	return nil
}

func (s *SocioeconomicRepositoryService) Get() error {
	return nil
}
