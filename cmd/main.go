package main

import (
	"cli-todo/internal/features/auth"
	todoHandler "cli-todo/internal/features/todo/handler"
	todoRepo "cli-todo/internal/features/todo/repository"
	todoService "cli-todo/internal/features/todo/service"
	"cli-todo/internal/httpserver"
)

func main() {
	repository := todoRepo.NewJsonFileHandler("storage/json/data.json")
	//sql := todoRepo.NewSqlRepo("storage/sqlLight/todo.db")
	service := todoService.NewService(repository)
	todoRouter := todoHandler.NewTodoHandler(service)

	googleAuthRouter := auth.NewGoogleAuthRouter()

	server := httpserver.NewServer(todoRouter, googleAuthRouter)
	server.StartServer()
}
