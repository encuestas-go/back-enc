package server

import (
	"github.com/encuestas-go/back-enc/internal/controllers"
	"github.com/labstack/echo/v4"
)

// ServerHandler handles the server logic outside of the package
type ServerHandler struct {
	ServerEcho     *echo.Echo
	RouterGroup    *echo.Group
	UserController *controllers.UserController
}

func InitServer() *ServerHandler {
	e := echo.New()
	userController := controllers.InitUserController()

	return &ServerHandler{
		ServerEcho:     e,
		UserController: userController,
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
