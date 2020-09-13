package server

import (
	"github.com/MrWebUzb/apiserver/app/handlers"
)

// ConfigureRouters ...
func (s *Server) ConfigureRouters() {
	s.HomeRouter()
	s.TodoRouter()
}

// HomeRouter ...
func (s *Server) HomeRouter() {
	homeHandler := handlers.NewHomeHandler()

	s.Router.Handle("/", homeHandler.Index()).Methods("GET")
}

// TodoRouter ...
func (s *Server) TodoRouter() {
	todoHandler := handlers.NewTodoHandler(
		s.Log,
		s.DB,
	)

	s.Router.Handle("/todo", todoHandler.Index()).Methods("GET")
	s.Router.Handle("/todo/{id:[0-9]+}", todoHandler.View()).Methods("GET")
}

// ConfigureErrorRouters ...
func (s *Server) ConfigureErrorRouters() {
	s.Router.NotFoundHandler = handlers.NotFoundHandler()
	s.Router.MethodNotAllowedHandler = handlers.MethodNotAllowedHandler()
}
