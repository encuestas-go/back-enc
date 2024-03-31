package main

import "github.com/encuestas-go/back-enc/internal/server"

func Build() {

	s := server.InitServer().
		StartActivityManagementRoutes().
		StartDemographicStatusRoutes().
		StartEconomicStatusRoutes().
		StartEventManagementRoutes().
		StartHouseholdInfrastructureRoutes().
		StartSatisfactorySurveysRoutes().
		StartServiceManagementRoutes().
		StartSocioeconomicStatusRoutes().
		StartTransportManagementRoutes().
		StartUserRoutes()

	s.StartServer()
}
