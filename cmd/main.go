package main

import (
	"cli-todo/internal/config"
	"cli-todo/internal/features/auth"
	"cli-todo/internal/features/todo"
	"cli-todo/internal/httpserver"
)

func main() {
	dbHandler := todo.NewJsonFileHandler("storage/json/data.json")
	service := todo.NewService(dbHandler)
	todoRouter := todo.NewTodoRouter(service)
	googleEnv := config.NewGoogleEnv()
	googleAuthRouter := auth.NewGoogleAuthRouter(googleEnv)

	server := httpserver.NewServer(todoRouter, googleAuthRouter)
	server.StartServer()
}
