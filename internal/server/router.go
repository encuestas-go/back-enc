package server

// StartRoutes initialize the routes to the group /api/v1
func (s *ServerHandler) StartRoutes() *ServerHandler {
	s.RouterGroup = s.ServerEcho.Group("/api/v1")
	return s
}

// InitalizeUserRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/usuario
//	/actualizar/usuario
//	/eliminar/usuario
//	/consultar/usuario
func (s *ServerHandler) InitalizeUserRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/usuario", s.UserController.CreateUser)
	s.RouterGroup.PUT("/actualizar/usuario", s.UserController.UpdateUser)
	s.RouterGroup.DELETE("/eliminar/usuario", s.UserController.DeleteUser)
	s.RouterGroup.GET("/consultar/usuario", s.UserController.GetUser)
	return s
}

// StartSocioeconomicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are
//
//	/crear/nivelSocioeconomico
//	/actualizar/nivelSocioeconomico
//	/eliminar/nivelSocioeconomico
//	/consultar/nivelSocioeconomico
func (s *ServerHandler) StartSocioeconomicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelSocioeconomico", s.SocioeconomicController.CreateSocioeconomicStatus)
	s.RouterGroup.PUT("/actualizar/nivelSocioeconomico", s.SocioeconomicController.UpdateSocioeconomicStatus)
	s.RouterGroup.DELETE("/eliminar/nivelSocioeconomico", s.SocioeconomicController.DeleteSocioeconomicStatus)
	s.RouterGroup.GET("/consultar/nivelSocioeconomico", s.SocioeconomicController.GetSocioeconomicStatus)
	return s
}

// StartEconomicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are
//
//	/crear/nivelEconomico
//	/actualizar/nivelEconomico
//	/eliminar/nivelEconomico
//	/consultar/nivelEconomico
func (s *ServerHandler) StartEconomicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelEconomico", s.EconomicController.CreateEconomicStatus)
	s.RouterGroup.PUT("/actualizar/nivelEconomico", s.EconomicController.UpdateEconomicStatus)
	s.RouterGroup.DELETE("/eliminar/nivelEconomico", s.EconomicController.DeleteEconomicStatus)
	s.RouterGroup.GET("/consultar/nivelEconomico", s.EconomicController.GetEconomicStatus)
	return s
}

// StartTransportManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are
//
//	/crear/medioTransporte
//	/actualizar/medioTransporte
//	/eliminar/medioTransporte
//	/consultar/medioTransporte
func (s *ServerHandler) StartTransportManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/medioTransporte", s.TransportController.CreateTransportManagement)
	s.RouterGroup.PUT("/actualizar/medioTransporte", s.TransportController.UpdateTransportManagement)
	s.RouterGroup.DELETE("/eliminar/medioTransporte", s.TransportController.DeleteTransportManagement)
	s.RouterGroup.GET("/consultar/medioTransporte", s.TransportController.GetTransportManagement)
	return s
}

// StartHouseholdInfrastructureRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are
//
//	/crear/InfraestructuraCasa
//	/actualizar/InfraestructuraCasa
//	/eliminar/InfraestructuraCasa
//	/consultar/InfraestructuraCasa
func (s *ServerHandler) StartHouseholdInfrastructureRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/InfraestructuraCasa", s.HouseInfrastructureController.CreateHouseInfrastructure)
	s.RouterGroup.PUT("/actualizar/InfraestructuraCasa", s.HouseInfrastructureController.UpdateHouseInfrastructure)
	s.RouterGroup.DELETE("/eliminar/InfraestructuraCasa", s.HouseInfrastructureController.DeleteHouseInfrastructure)
	s.RouterGroup.GET("/consultar/InfraestructuraCasa", s.HouseInfrastructureController.GetHouseInfrastructure)
	return s
}

// StartDemographicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/nivelDemografico
//	/actualizar/nivelDemografico
//	/eliminar/nivelDemografico
//	/consultar/nivelDemografico
func (s *ServerHandler) StartDemographicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelDemografico", s.DemographicController.CreateDemographicStatus)
	s.RouterGroup.PUT("/actualizar/nivelDemografico", s.DemographicController.UpdateDemographicStatus)
	s.RouterGroup.DELETE("/eliminar/nivelDemografico", s.DemographicController.DeleteDemographicStatus)
	s.RouterGroup.GET("/consultar/nivelDemografico", s.DemographicController.GetDemographicStatus)
	return s
}

// StartActivitiesManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/actividad
//	/actualizar/actividad
//	/eliminar/actividad
//	/consultar/actividad
func (s *ServerHandler) StartActivityManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/actividad", s.ActivityController.CreateActivityManagement)
	s.RouterGroup.PUT("/actualizar/actividad", s.ActivityController.UpdateActivityManagement)
	s.RouterGroup.DELETE("/eliminar/actividad", s.ActivityController.DeleteActivityManagement)
	s.RouterGroup.GET("/consultar/actividad", s.ActivityController.GetActivityManagement)
	return s
}

// StartActivitiesManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/servicio
//	/actualizar/servicio
//	/eliminar/servicio
//	/consultar/servicio
/*
func (s *ServerHandler) StartServiceManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/servicio")
	s.RouterGroup.PUT("/actualizar/servicio")
	s.RouterGroup.DELETE("/eliminar/servicio")
	s.RouterGroup.GET("/consultar/servicio")
	return s
}

func (s *ServerHandler) StartEventManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/evento",)
	s.RouterGroup.PUT("/actualizar/evento",)
	s.RouterGroup.DELETE("/eliminar/evento",)
	s.RouterGroup.GET("/consultar/evento",)
	return s
}

func (s *ServerHandler) StartSatisfactorySurveysRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/encuestaSatisfaccion",)
	s.RouterGroup.PUT("/actualizar/encuestaSatisfaccion",)
	s.RouterGroup.DELETE("/eliminar/encuestaSatisfaccion",)
	s.RouterGroup.GET("/consultar/encuestaSatisfaccion",)
	return s
}
*/
