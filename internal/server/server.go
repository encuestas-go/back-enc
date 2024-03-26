package server

// Agregar imports

type ServerHandler struct {
	// aqui agregar el tipo de dato del servidor
	ServerEcho int
}

func InitServer() *ServerHandler  {
	e := echo.New()
	return &ServerHandler{
		ServerEcho: e
	}
}

func (s *ServerHandler) StartServer() {
	s.ServerEcho.Logger.Fatal(s.ServerEcho.Start(":3000"))
}
