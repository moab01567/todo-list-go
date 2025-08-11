package main

import (
	"cli-todo/internal/config"
	"cli-todo/internal/features/auth"
	"cli-todo/internal/features/todo"
	"cli-todo/internal/httpserver"
	"os"
)

func main() {
	dbHandler := todo.NewJsonFileHandler("storage/json/data.json")
	service := todo.NewService(dbHandler)
	todoRouter := todo.NewTodoRouter(service)
	googleAuthRouter := auth.NewGoogleAuthRouter(
		os.Getenv(config.GOOGLE_AUTH_URL),
		os.Getenv(config.GOOGLE_SCOPE),
		os.Getenv(config.GOOGLE_REDIRECT_URL),
		os.Getenv(config.GOOGLE_CLIENT_ID))

	server := httpserver.NewServer(todoRouter, googleAuthRouter)
	server.StartServer()
}
