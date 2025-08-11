package httpserver

import (
	"log"
	"net/http"
)

type Router interface {
	GetHandler() http.Handler
}

type Server struct {
	todoRouter       Router
	authGoogleRouter Router
}

func NewServer(todoRouter, authGoogleRouter Router) *Server {
	return &Server{todoRouter: todoRouter, authGoogleRouter: authGoogleRouter}
}

func (server *Server) StartServer() {

	log.Fatal(http.ListenAndServe(":8080", server.authGoogleRouter.GetHandler()))
}
