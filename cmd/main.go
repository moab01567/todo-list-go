package main

import (
	"cli-todo/internal/features/googleauth"
	todoHandler "cli-todo/internal/features/todo/handler"
	todoRepo "cli-todo/internal/features/todo/repository"
	todoService "cli-todo/internal/features/todo/service"
	"cli-todo/internal/httpserver"
)

func main() {
	repository := todoRepo.NewJsonFileHandler("storage/json/data.json")
	//sql := todoRepo.NewSqlRepo("storage/sqlLight/db.db")
	service := todoService.NewService(repository)
	todoRouter := todoHandler.NewTodoHandler(service)

	// init google googleauth :)
	googleEnv := googleauth.NewGoogleEnv()
	googleAuthRouter := googleauth.NewGoogleHandler(googleEnv)

	server := httpserver.NewServer(todoRouter, googleAuthRouter)
	server.StartServer()
}
