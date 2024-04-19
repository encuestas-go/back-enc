package repository

import "github.com/encuestas-go/back-enc/internal/database"

type RepositoryStorage struct {
	UserRespository         *UserRepositoryService
	SocioeconomicRepository *SocioeconomicRepositoryService
}

func GetRepository() *RepositoryStorage {
	db := database.ConnectToDB()

	userRespositoryStorage := InitializeUserRepository(db)
	socioeconomicRepositoryStorage := InitializeSocioeconomicRepository(db)

	return &RepositoryStorage{
		UserRespository:         userRespositoryStorage,
		SocioeconomicRepository: socioeconomicRepositoryStorage,
	}
}
