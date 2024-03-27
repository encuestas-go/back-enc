package server

// StartRoutes initialize the routes to the group /api/v1
func (s *ServerHandler) StartRoutes() *ServerHandler {
	s.RouterGroup = s.ServerEcho.Group("/api/v1")
	return s
}

// InitalizeUserRoutes creates the routes for the user requirement based on the group: /v1/api.
// The routes are:
// - /crear/usuario
func (s *ServerHandler) InitalizeUserRoutes() *ServerHandler {
	s.RouterGroup.POST("/crear/usuario", s.UserController.CreateUser)
	s.RouterGroup.PUT("/actualizar/usuario", nil)
	s.RouterGroup.DELETE("/eliminar/usuario", nil)

	return s
}
