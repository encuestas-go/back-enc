package server

import (
	"github.com/labstack/echo/v4"
)

type ServerHandler struct {
	ServerEcho *echo.Echo
}

func InitServer() *ServerHandler {
	e := echo.New()
	return &ServerHandler{
		ServerEcho: e,
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
