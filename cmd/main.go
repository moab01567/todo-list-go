package main

import (
	"cli-todo/internal/httpserver"
	"cli-todo/internal/todo"
)

func main() {
	dbHandler := todo.NewJsonFileHandler("storage/json/data.json")
	service := todo.NewService(dbHandler)
	todoRouter := todo.NewTodoRouter(service)

	server := httpserver.NewServer(todoRouter)
	server.StartServer()
}
