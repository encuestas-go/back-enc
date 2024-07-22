package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type errorUnauthorizedMessage struct {
	Message string `json:"message,omitempty"`
}

// StartRouterGroup StartRoutes initialize the routes to the group /api/v1
func (s *ServerHandler) StartRouterGroup() *ServerHandler {
	s.RouterGroup = s.ServerEcho.Group("/api/v1")

	s.RouterGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			cookieIDUser, err := c.Cookie("id_user")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
					Message: "Unable to retrieve ID User cookie",
				})
			}

			IDUserConverted, err := strconv.Atoi(cookieIDUser.Value)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
					Message: "ID User cannot be converted",
				})
			}

			cookieIDTypeUser, err := c.Cookie("id_type_user")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
					Message: "Unable to retrieve ID Type User cookie",
				})
			}

			IDTypeUserConverted, err := strconv.Atoi(cookieIDTypeUser.Value)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
					Message: "ID Type User cannot be converted",
				})
			}

			if IDUserConverted == 0 || IDTypeUserConverted == 0 {
				return c.JSON(http.StatusUnauthorized, errorUnauthorizedMessage{
					Message: "Invalid ID's values for cookie information",
				})
			}

			return next(c)
		}
	})
	return s
}

// StartUserRoutes 1.-InitalizeUserRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/usuario
//	/actualizar/usuario
//	/eliminar/usuario
//	/consultar/usuario
func (s *ServerHandler) StartUserRoutes() *ServerHandler {
	s.ServerEcho.POST("/crear/usuario", s.GenericController.UserController.Create)
	s.RouterGroup.PUT("/actualizar/usuario", s.GenericController.UserController.Update)
	s.RouterGroup.DELETE("/eliminar/usuario", s.GenericController.UserController.Delete)
	s.RouterGroup.GET("/consultar/usuario", s.GenericController.UserController.Get)

	// login
	s.ServerEcho.POST("/login", s.GenericController.UserController.Login)

	// restore password
	s.ServerEcho.POST("/login/reset-password", s.GenericController.UserController.ResetPassword)

	return s
}

// StartSocioeconomicStatusRoutes 2.-StartSocioeconomicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
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

// StartEconomicStatusRoutes 3.-StartEconomicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
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

	s.RouterGroup.GET("/reporte/SituacionActualEstudiante", s.GenericController.EconomicStatusController.GetStudentSituationReport)

	return s
}

// StartTransportManagementRoutes 4.- StartTransportManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
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

	s.RouterGroup.GET("/reporte/TransportePrincipal", s.GenericController.TransportController.GetTransportReport)

	return s
}

// StartHouseholdInfrastructureRoutes 5.-StartHouseholdInfrastructureRoutes creates the routes for the user requirement based on the group: /v1/api.
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

// StartDemographicStatusRoutes 6.-StartDemographicStatusRoutes creates the routes for the user requirement based on the group: /v1/api.
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

	s.RouterGroup.GET("/reporte/IngresosMensuales", s.GenericController.DemographicController.GetAllIncomeAmountReport)
	s.RouterGroup.GET("/reporte/TipoViviendaCondicion", s.GenericController.DemographicController.GetHouseTypeConditionReport)
	return s
}

// StartActivityManagementRoutes 7.-StartActivitiesManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
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

	s.RouterGroup.GET("/reporte/PreferenciasActividades", s.GenericController.ActivityManagementController.GetCulturalActivitiesReports)

	return s
}

// StartServiceManagementRoutes 8.-StartActivitiesManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
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

	s.RouterGroup.GET("/reporte/ProveedorInternetReporte", s.GenericController.ServiceManagementController.GetAllInternetProvidersReport)

	return s
}

// StartEventManagementRoutes 9.- StartEventManagementRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/evento
//	/actualizar/evento
//	/eliminar/evento
//	/consultar/evento
func (s *ServerHandler) StartEventManagementRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/evento", s.GenericController.EventManagementController.Create, NeedAdmin)
	s.RouterGroup.PUT("/actualizar/evento", s.GenericController.EventManagementController.Update, NeedAdmin)
	s.RouterGroup.DELETE("/eliminar/evento", s.GenericController.EventManagementController.Delete, NeedAdmin)
	s.RouterGroup.GET("/consultar/evento", s.GenericController.EventManagementController.Get)

	return s
}

// StartSatisfactorySurveysRoutes 10.- StartSatisfactorySurveysRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
//
//	/crear/encuestaSatisfaccion
//	/actualizar/encuestaSatisfaccion
//	/eliminar/encuestaSatisfaccion
//	/consultar/encuestaSatisfaccion
func (s *ServerHandler) StartSatisfactorySurveysRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Create)
	s.RouterGroup.POST("/crear/encuestaLikert", s.GenericController.SatisfactorySurveyController.CreateLikertSurvey)
	s.RouterGroup.PUT("/actualizar/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Update)
	s.RouterGroup.DELETE("/eliminar/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Delete)
	s.RouterGroup.GET("/consultar/encuestaSatisfaccion", s.GenericController.SatisfactorySurveyController.Get)
	s.RouterGroup.GET("/consultar/encuestaHorario", s.GenericController.SatisfactorySurveyController.GetSchedule)

	return s
}

// StartForumRoutes  11.-StartForumRoutes creates de routes for the forum, to publish a question by part of an user,
// and answers from other user.
// The routes are:
//
//	/ingresar/foro
func (s *ServerHandler) StartForumRoutes() *ServerHandler {
	s.RouterGroup.POST("/publicar/preguntaForo", s.GenericController.ForumController.CreateQuestion)
	s.RouterGroup.POST("/publicar/respuestaForo", s.GenericController.ForumController.CreateAnswer)
	s.RouterGroup.GET("/obtener/datosForo", s.GenericController.ForumController.Get)

	return s
}

// StartForumRoutes  12.-StartGetAddressRoutes creates the route for the map, to get all users addresses
// The routes are:
//
//	/direccionMapa
func (s *ServerHandler) StartMapRoutes() *ServerHandler {
	s.RouterGroup.GET("/direccionMapa", s.GenericController.MapAddressController.Get)

	return s
}

// StarBackupRoutes  13.-StartBackupRoutes creates the routes for the backup of information from database, using a date format YYYY-MM-DD to find
// The routes are:
//
//	/crearBackups
//	/obtenerBackups
func (s *ServerHandler) StartBackupRoutes() *ServerHandler {
	s.RouterGroup.POST("/crearBackups", s.GenericController.BackupController.Create)
	s.RouterGroup.GET("/obtenerBackups", s.GenericController.BackupController.Get)
	s.RouterGroup.POST("/ejectuarBackup", s.GenericController.BackupController.Restore)

	return s
}
