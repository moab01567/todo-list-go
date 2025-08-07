package httpserver

import (
	"cli-todo/internal/todo"
	"log"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/create", todo.CreateTodoRouter())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
