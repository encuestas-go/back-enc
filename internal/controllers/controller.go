package controllers

import (
	"database/sql"

	"github.com/encuestas-go/back-enc/internal/repository"
)

type ControllerMessageResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type GenericController struct {
	ActivityManagementController  *ActivityManagementController
	DemographicController         *DemographicStatusController
	EconomicStatusController      *EconomicStatusController
	EventManagementController     *EventManagementController
	HouseInfrastructureController *HouseInfrastructureController
	SatisfactorySurveyController  *SatisfactorySurveyController
	ServiceManagementController   *ServiceManagementController
	SocioeconomicStatusController *SocioeconomicStatusController
	TransportController           *TransportController
	UserController                *UserController
	ForumController               *ForumController
	MapAddressController          *MapAddressController
}

func InitGenericController(db *sql.DB) *GenericController {
	repositories := repository.GetRepository(db)

	return &GenericController{
		ActivityManagementController:  InitActivityController(repositories.CulturalActivityRepository),
		DemographicController:         InitDemographicController(repositories.DemographicRepository),
		EconomicStatusController:      InitEconomicController(repositories.EconomicRepository),
		EventManagementController:     InitEventManagementController(repositories.EventRepository),
		HouseInfrastructureController: InitHouseInfrastructureController(repositories.InfrastructureRepository),
		SatisfactorySurveyController:  InitSatisfactorySurveyController(),
		ServiceManagementController:   InitServiceManagementController(repositories.ServicesRepository),
		SocioeconomicStatusController: InitSocioeconomicController(repositories.SocioeconomicRepository),
		TransportController:           InitTransportController(repositories.TransportRepository),
		UserController:                InitUserController(repositories.UserRespository),
		ForumController:               InitForumController(repositories.ForumRepository),
		MapAddressController:          InitMapController(repositories.MapRepository),
	}
}
