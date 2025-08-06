package main

import (
	"cli-todo/internal/todo"
	"fmt"
)

func main() {
	//CmdManager(os.Args[0], os.Args[1:]
	var jsonFileHandler todo.DbHandler = &todo.JsonFileHandler{FilePath: "storage/json/data.json"}
	var service todo.Service = todo.Service{DbHandler: jsonFileHandler}
	err := service.DeleteTodo(46)
	fmt.Println(err)

}
