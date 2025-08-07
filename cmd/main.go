package main

import "cli-todo/internal/httpserver"

func main() {
	//var jsonFileHandler todo.DbHandler = &todo.JsonFileHandler{FilePath: "storage/json/data.json"}
	//var service *todo.Service = &todo.Service{DbHandler: jsonFileHandler}
	httpserver.StartServer()
}
