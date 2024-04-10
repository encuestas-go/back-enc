package repository

import (
	"database/sql"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type UserRepositoryService struct {
	db *sql.DB
}

func InitializeUserRepository(db *sql.DB) *UserRepositoryService {
	return &UserRepositoryService{
		db: db,
	}
}

func (u *UserRepositoryService) Insert(user domain.User) error {
	result, err := u.db.Exec(`
	`)
	///.....
	///....
	//.....

	return nil
}
