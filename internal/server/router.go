package server

// StartRoutes initialize the routes to the group /api/v1
func (s *ServerHandler) StartRoutes() *ServerHandler {
	s.RouterGroup = s.ServerEcho.Group("/api/v1")
	return s
}

// InitalizeUserRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//   - /crear/usuario
//     /actualizar/usuario
//     /eliminar/usuario
//     /consultar/usuario
func (s *ServerHandler) InitalizeUserRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/usuario", s.UserController.CreateUser)
	s.RouterGroup.PUT("/actualizar/usuario", s.UserController.UpdateUser)
	s.RouterGroup.DELETE("/eliminar/usuario", s.UserController.DeleteUser)
	s.RouterGroup.GET("/consultar/usuario", s.UserController.GetUser)
	return s
}

/*
func (s *ServerHandler) StartSocioeconomicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelSocioeconomico",)
	s.RouterGroup.PUT("/actualizar/nivelSocioeconomico",)
	s.RouterGroup.DELETE("/eliminar/nivelSocioeconomico",)
	s.RouterGroup.GET("/consultar/nivelSocioeconomico",)
	return s
}

func (s *ServerHandler) StartEconomicStatusRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelEconomico",)
	s.RouterGroup.PUT("/crear/nivelEconomico",)
	s.RouterGroup.DELETE("/crear/nivelEconomico",)
	s.RouterGroup.GET("/crear/nivelEconomico",)
	return s
}

func (s *ServerHandler) StartTransportManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/medioTransporte",)
	s.RouterGroup.PUT("/actualizar/medioTransporte",)
	s.RouterGroup.DELETE("/eliminar/medioTransporte",)
	s.RouterGroup.GET("/consultar/medioTransporte",)
	return s
}

func (s *ServerHandler) StartHouseholdInfrastructureRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/InfraestructuraCasa",)
	s.RouterGroup.PUT("/actualizar/InfraestructuraCasa",)
	s.RouterGroup.DELETE("/eliminar/InfraestructuraCasa",)
	s.RouterGroup.GET("/consultar/InfraestructuraCasa",)
	return s
}

func (s *ServerHandler) StartDemographicManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/nivelDemografico",)
	s.RouterGroup.PUT("/actualizar/nivelDemografico",)
	s.RouterGroup.DELETE("/eliminar/nivelDemografico",)
	s.RouterGroup.GET("/consultar/nivelDemografico",)
	return s
}

func (s *ServerHandler) StartActivitiesManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/actividad",)
	s.RouterGroup.PUT("/actualizar/actividad",)
	s.RouterGroup.DELETE("/eliminar/actividad",)
	s.RouterGroup.GET("/consultar/actividad",)
	return s
}

func (s *ServerHandler) StartServiceManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/servicio",)
	s.RouterGroup.PUT("/actualizar/servicio",)
	s.RouterGroup.DELETE("/eliminar/servicio",)
	s.RouterGroup.GET("/consultar/servicio",)
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
