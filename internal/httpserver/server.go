package httpserver

import (
	"log"
	"net/http"
)

type Handler interface {
	Routes() http.Handler
}

type Server struct {
	todoRouter       Handler
	authGoogleRouter Handler
}

func NewServer(todoRouter, authGoogleRouter Handler) *Server {
	return &Server{todoRouter: todoRouter, authGoogleRouter: authGoogleRouter}
}

func (server *Server) StartServer() {

	log.Fatal(http.ListenAndServe(":8080", server.authGoogleRouter.Routes()))
}
