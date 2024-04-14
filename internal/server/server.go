package server

import (
	"github.com/encuestas-go/back-enc/internal/controllers"
	"github.com/labstack/echo/v4"
)

// ServerHandler handles the server logic outside of the package
type ServerHandler struct {
	ServerEcho        *echo.Echo
	RouterGroup       *echo.Group
	GenericController *controllers.GenericController
}

func InitServer() *ServerHandler {
	e := echo.New()
	return &ServerHandler{
		ServerEcho:        e,
		GenericController: controllers.InitGenericController(),
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
