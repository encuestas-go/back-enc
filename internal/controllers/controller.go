package controllers

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
}

func InitGenericController() *GenericController {
	return &GenericController{
		ActivityManagementController:  InitActivityController(),
		DemographicController:         InitDemographicController(),
		EconomicStatusController:      InitEconomicController(),
		EventManagementController:     InitEventManagementController(),
		HouseInfrastructureController: InitHouseInfrastructureController(),
		SatisfactorySurveyController:  InitSatisfactorySurveyController(),
		ServiceManagementController:   InitServiceManagementController(),
		SocioeconomicStatusController: InitSocioeconomicController(),
		TransportController:           InitTransportController(),
		UserController:                InitUserController(),
	}
}
