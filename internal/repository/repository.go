package repository

import (
	"database/sql"
)

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
	ForumRepository            *ForumRepositoryService
	MapRepository              *MapRepositoryService
}

func GetRepository(db *sql.DB) *RepositoryStorage {
	userRespositoryStorage := InitializeUserRepository(db)
	socioeconomicRepositoryStorage := InitializeSocioeconomicRepository(db)
	economicRepositoryStorage := InitializeEconomicRepository(db)
	transportRepositoryStorage := InitializeTransportRepository(db)
	infrastructureRepositoryStorage := InitializeInfrastructureRepository(db)
	demographicRepositoryStorage := InitializeDemographicRepository(db)
	activityRepositoryStorage := InitializeCulturalActivityRepository(db)
	servicesRepositoryStorage := InitializeServicesRepository(db)
	eventRepositoryStorage := InitializeEventRepository(db)
	forumRepositoryStorage := InitializeForumRepository(db)
	mapRepositoryStorage := InitializeMapRepository(db)

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
		ForumRepository:            forumRepositoryStorage,
		MapRepository:              mapRepositoryStorage,
	}
}
