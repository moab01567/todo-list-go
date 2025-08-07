package httpserver

import (
	"cli-todo/internal/todo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	greetings := func(w http.ResponseWriter, r *http.Request) {

		var header http.Header = w.Header()
		header.Set("Content-Type", "application/json")
		var todos []todo.Todo
		todos = append(todos, todo.CreateRandomTodo())
		dataBytes, _ := json.Marshal(todos)
		bytesSend, _ := w.Write(dataBytes)
		fmt.Println("total bytes send:", bytesSend)

	}
	http.HandleFunc("/", greetings)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
