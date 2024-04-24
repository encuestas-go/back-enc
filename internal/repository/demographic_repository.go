package repository

import "database/sql"

type DemographicRepositoryService struct {
	db *sql.DB
}

func InitializeDemographicRepository(db *sql.DB) *DemographicRepositoryService {
	return &DemographicRepositoryService{
		db: db,
	}
}
