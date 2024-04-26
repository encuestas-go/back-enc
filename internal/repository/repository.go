package repository

import "github.com/encuestas-go/back-enc/internal/database"

type RepositoryStorage struct {
	UserRespository            *UserRepositoryService
	SocioeconomicRepository    *SocioeconomicRepositoryService
	EconomicRepository         *EconomicRepositoryService
	TransportRepository        *TransportRespositoryService
	InfrastructureRepository   *HouseInfrastructureRepositoryService
	DemographicRepository      *DemographicRepositoryService
	CulturalActivityRepository *CulturalActivityRepositoryService
	ServicesRepository         *ServicesRepositoryService
	EventRepository            *EventRepositoryService
}

func GetRepository() *RepositoryStorage {
	db := database.ConnectToDB()

	userRespositoryStorage := InitializeUserRepository(db)
	socioeconomicRepositoryStorage := InitializeSocioeconomicRepository(db)
	economicRepositoryStorage := InitializeEconomicRepository(db)
	transportRepositoryStorage := InitializeTransportRepository(db)
	infrastructureRepositoryStorage := InitializeInfrastructureRepository(db)
	demographicRepositoryStorage := InitializeDemographicRepository(db)
	activityRepositoryStorage := InitializeCulturalActivityRepository(db)
	servicesRepositoryStorage := InitializeServicesRepository(db)
	eventRepositoryStorage := InitializeEventRepository(db)

	return &RepositoryStorage{
		UserRespository:            userRespositoryStorage,
		SocioeconomicRepository:    socioeconomicRepositoryStorage,
		EconomicRepository:         economicRepositoryStorage,
		TransportRepository:        transportRepositoryStorage,
		InfrastructureRepository:   infrastructureRepositoryStorage,
		DemographicRepository:      demographicRepositoryStorage,
		CulturalActivityRepository: activityRepositoryStorage,
		ServicesRepository:         servicesRepositoryStorage,
		EventRepository:            eventRepositoryStorage,
	}
}
