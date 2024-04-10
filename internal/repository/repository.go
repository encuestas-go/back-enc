package repository

import "github.com/encuestas-go/back-enc/internal/database"

type RepositoryStorage struct {
	UserRespository *UserRepositoryService
}

func GetRepository() *RepositoryStorage {
	db := database.ConnectToDB()

	userRespositoryStorage := InitializeUserRepository(db)

	return &RepositoryStorage{
		UserRespository: userRespositoryStorage,
	}
}
