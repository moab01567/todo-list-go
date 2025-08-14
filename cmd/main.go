package main

import (
	"cli-todo/internal/config"
	"cli-todo/internal/features/auth"
	"cli-todo/internal/features/todo"
	"cli-todo/internal/httpserver"
)

func main() {
	repository := todo.NewJsonFileHandler("storage/json/data.json")
	service := todo.NewService(repository)
	todoRouter := todo.NewTodoRouter(service)
	googleEnv := config.NewGoogleEnv()
	googleAuthRouter := auth.NewGoogleAuthRouter(googleEnv)

	server := httpserver.NewServer(todoRouter, googleAuthRouter)
	server.StartServer()
}
