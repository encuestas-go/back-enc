package server

import (
	"github.com/encuestas-go/back-enc/internal/controllers"
	"github.com/labstack/echo/v4"
)

// ServerHandler handles the server logic outside of the package
type ServerHandler struct {
	ServerEcho                    *echo.Echo
	RouterGroup                   *echo.Group
	UserController                *controllers.UserController
	SocioeconomicController       *controllers.SocioeconomicController
	EconomicController            *controllers.EconomicController
	TransportController           *controllers.TransportController
	HouseInfrastructureController *controllers.HouseInfrastructure
	DemographicController         *controllers.DemographicController
	ActivityController            *controllers.ActivityController
}

func InitServer() *ServerHandler {
	e := echo.New()
	userController := controllers.InitUserController()
	socioeconomicController := controllers.InitSocioeconomicController()
	economicController := controllers.InitEconomicController()
	transportController := controllers.InitTransportController()
	houseInfrastructure := controllers.InitHouseInfrastructureController()
	demographicController := controllers.InitDemographicController()
	activityController := controllers.InitActivityController()

	return &ServerHandler{
		ServerEcho:                    e,
		UserController:                userController,
		SocioeconomicController:       socioeconomicController,
		EconomicController:            economicController,
		TransportController:           transportController,
		HouseInfrastructureController: houseInfrastructure,
		DemographicController:         demographicController,
		ActivityController:            activityController,
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
