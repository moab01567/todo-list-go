package httpserver

import (
	"log"
	"net/http"
)

type Router interface {
	GetHandler() http.Handler
}

type Server struct {
	todoRouter Router
}

func NewServer(todoRouter Router) *Server {
	return &Server{todoRouter: todoRouter}
}

func (server *Server) StartServer() {

	log.Fatal(http.ListenAndServe(":8080", server.todoRouter.GetHandler()))
}
