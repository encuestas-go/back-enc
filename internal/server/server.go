package server

import (
	"github.com/encuestas-go/back-enc/internal/controllers"
	"github.com/encuestas-go/back-enc/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ServerHandler handles the server logic outside of the package
type ServerHandler struct {
	ServerEcho        *echo.Echo
	RouterGroup       *echo.Group
	GenericController *controllers.GenericController
}

func InitServer() *ServerHandler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:5501", "http://localhost:3000", "http://localhost:5501"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST, echo.HEAD, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
	}))

	db := database.ConnectToDB()

	return &ServerHandler{
		ServerEcho:        e,
		GenericController: controllers.InitGenericController(db),
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
