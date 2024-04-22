package repository

import "github.com/encuestas-go/back-enc/internal/database"

type RepositoryStorage struct {
	UserRespository         *UserRepositoryService
	SocioeconomicRepository *SocioeconomicRepositoryService
	EconomicRepository      *EconomicRepositoryService
	TransportRepository     *TransportRespositoryService
}

func GetRepository() *RepositoryStorage {
	db := database.ConnectToDB()

	userRespositoryStorage := InitializeUserRepository(db)
	socioeconomicRepositoryStorage := InitializeSocioeconomicRepository(db)
	economicRepositoryStorage := InitializeEconomicRepository(db)
	transportRepositoryStorage := InitializeTransportRepository(db)

	return &RepositoryStorage{
		UserRespository:         userRespositoryStorage,
		SocioeconomicRepository: socioeconomicRepositoryStorage,
		EconomicRepository:      economicRepositoryStorage,
		TransportRepository:     transportRepositoryStorage,
	}
}
