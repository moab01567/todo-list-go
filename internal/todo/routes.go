package todo

import "net/http"

func CreateTodoRouter() func(w http.ResponseWriter, r *http.Request) {
	service := NewService(NewJsonFileHandler("storage/json/data.json"))
	return func(w http.ResponseWriter, r *http.Request) {
		service.AddTodo("test")
		w.Write([]byte("created"))
	}
}
