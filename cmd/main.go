package main

import (
	"cli-todo/internal/config"
	"cli-todo/internal/features/auth"
	todoHandler "cli-todo/internal/features/todo/handler"
	todoRepo "cli-todo/internal/features/todo/repository"
	todoService "cli-todo/internal/features/todo/service"
	"cli-todo/internal/httpserver"
)

func main() {
	repository := todoRepo.NewJsonFileHandler("storage/json/data.json")
	service := todoService.NewService(repository)
	todoRouter := todoHandler.NewTodoHandler(service)

	googleEnv := config.NewGoogleEnv()
	googleAuthRouter := auth.NewGoogleAuthRouter(googleEnv)

	server := httpserver.NewServer(todoRouter, googleAuthRouter)
	server.StartServer()
}
