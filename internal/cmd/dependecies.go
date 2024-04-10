package main

import "github.com/encuestas-go/back-enc/internal/server"

func build() {
	// steps:
	//  start server echo
	// start the group of the routes: v1/api
	// start each routes by service ex: user routes -> /crear/usuario
	// start other routes needed
	s := server.InitServer().StartRouterGroup().StartUserRoutes()

	s.StartServer()
}
