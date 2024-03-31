package server

// StartRoutes initialize the routes to the group /api/v1
func (s *ServerHandler) StartRoutes() *ServerHandler {
	s.RouterGroup = s.ServerEcho.Group("/api/v1")
	return s
}

// 1.-InitalizeUserRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/usuario
//	/actualizar/usuario
//	/eliminar/usuario
//	/consultar/usuario
func (s *ServerHandler) InitalizeUserRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/usuario", s.GenericController.UserController.Create)
	s.RouterGroup.PUT("/actualizar/usuario", s.GenericController.UserController.Update)
	s.RouterGroup.DELETE("/eliminar/usuario", s.GenericController.UserController.Delete)
	s.RouterGroup.GET("/consultar/usuario", s.GenericController.UserController.Get)

	s.RouterGroup.POST("/login", nil)
	s.RouterGroup.POST("/logOut", nil)

	return s
}

// 2.-StartSocioeconomicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/nivelSocioeconomico
//	/actualizar/nivelSocioeconomico
//	/eliminar/nivelSocioeconomico
//	/consultar/nivelSocioeconomico
func (s *ServerHandler) StartSocioeconomicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelSocioeconomico", s.GenericController.SocioeconomicStatusController.Create)
	s.RouterGroup.PUT("/actualizar/nivelSocioeconomico", s.GenericController.SocioeconomicStatusController.Update)
	s.RouterGroup.DELETE("/eliminar/nivelSocioeconomico", s.GenericController.SocioeconomicStatusController.Delete)
	s.RouterGroup.GET("/consultar/nivelSocioeconomico", s.GenericController.SocioeconomicStatusController.Get)

	return s
}

// 3.-StartEconomicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/nivelEconomico
//	/actualizar/nivelEconomico
//	/eliminar/nivelEconomico
//	/consultar/nivelEconomico
func (s *ServerHandler) StartEconomicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelEconomico", s.GenericController.EconomicStatusController.Create)
	s.RouterGroup.PUT("/actualizar/nivelEconomico", s.GenericController.EconomicStatusController.Update)
	s.RouterGroup.DELETE("/eliminar/nivelEconomico", s.GenericController.EconomicStatusController.Delete)
	s.RouterGroup.GET("/consultar/nivelEconomico", s.GenericController.EconomicStatusController.Get)

	return s
}

// 4.- StartTransportManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/medioTransporte
//	/actualizar/medioTransporte
//	/eliminar/medioTransporte
//	/consultar/medioTransporte
func (s *ServerHandler) StartTransportManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/medioTransporte", s.GenericController.TransportController.Create)
	s.RouterGroup.PUT("/actualizar/medioTransporte", s.GenericController.TransportController.Update)
	s.RouterGroup.DELETE("/eliminar/medioTransporte", s.GenericController.TransportController.Delete)
	s.RouterGroup.GET("/consultar/medioTransporte", s.GenericController.TransportController.Get)

	return s
}

// 5.-StartHouseholdInfrastructureRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/InfraestructuraCasa
//	/actualizar/InfraestructuraCasa
//	/eliminar/InfraestructuraCasa
//	/consultar/InfraestructuraCasa
func (s *ServerHandler) StartHouseholdInfrastructureRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/InfraestructuraCasa", s.GenericController.HouseInfrastructureController.Create)
	s.RouterGroup.PUT("/actualizar/InfraestructuraCasa", s.GenericController.HouseInfrastructureController.Update)
	s.RouterGroup.DELETE("/eliminar/InfraestructuraCasa", s.GenericController.HouseInfrastructureController.Delete)
	s.RouterGroup.GET("/consultar/InfraestructuraCasa", s.GenericController.HouseInfrastructureController.Get)

	return s
}

// 6.-StartDemographicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/nivelDemografico
//	/actualizar/nivelDemografico
//	/eliminar/nivelDemografico
//	/consultar/nivelDemografico
func (s *ServerHandler) StartDemographicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelDemografico", s.GenericController.DemographicController.Create)
	s.RouterGroup.PUT("/actualizar/nivelDemografico", s.GenericController.DemographicController.Update)
	s.RouterGroup.DELETE("/eliminar/nivelDemografico", s.GenericController.DemographicController.Delete)
	s.RouterGroup.GET("/consultar/nivelDemografico", s.GenericController.DemographicController.Get)

	return s
}

// 7.-StartActivitiesManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/actividad
//	/actualizar/actividad
//	/eliminar/actividad
//	/consultar/actividad
func (s *ServerHandler) StartActivityManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/actividad", s.GenericController.ActivityManagementController.Create)
	s.RouterGroup.PUT("/actualizar/actividad", s.GenericController.ActivityManagementController.Update)
	s.RouterGroup.DELETE("/eliminar/actividad", s.GenericController.ActivityManagementController.Delete)
	s.RouterGroup.GET("/consultar/actividad", s.GenericController.ActivityManagementController.Get)

	return s
}

// 8.-StartActivitiesManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/servicio
//	/actualizar/servicio
//	/eliminar/servicio
//	/consultar/servicio
func (s *ServerHandler) StartServiceManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/servicio", s.GenericController.ServiceManagementController.Create)
	s.RouterGroup.PUT("/actualizar/servicio", s.GenericController.ServiceManagementController.Update)
	s.RouterGroup.DELETE("/eliminar/servicio", s.GenericController.ServiceManagementController.Delete)
	s.RouterGroup.GET("/consultar/servicio", s.GenericController.ServiceManagementController.Get)

	return s
}

// 9.- StartEventManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/evento
//	/actualizar/evento
//	/eliminar/evento
//	/consultar/evento
func (s *ServerHandler) StartEventManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/evento", s.GenericController.EventManagementController.Create)
	s.RouterGroup.PUT("/actualizar/evento", s.GenericController.EventManagementController.Update)
	s.RouterGroup.DELETE("/eliminar/evento", s.GenericController.EventManagementController.Delete)
	s.RouterGroup.GET("/consultar/evento", s.GenericController.EconomicStatusController.Get)

	return s
}

// 10.- StartSatisfactorySurveysRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/encuestaSatisfaccion
//	/actualizar/encuestaSatisfaccion
//	/eliminar/encuestaSatisfaccion
//	/consultar/encuestaSatisfaccion
func (s *ServerHandler) StartSatisfactorySurveysRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Create)
	s.RouterGroup.PUT("/actualizar/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Update)
	s.RouterGroup.DELETE("/eliminar/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Delete)
	s.RouterGroup.GET("/consultar/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Get)

	return s
}
