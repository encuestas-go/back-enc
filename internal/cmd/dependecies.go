package main

import "github.com/encuestas-go/back-enc/internal/server"

func Build() {

	s := server.InitServer()

	s.StartServer()
}
