package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Service interface {
	AddTodo(name string) error
	DeleteTodo(id string) error
	GetTodos() ([]Todo, error)
}

type TodoRouter struct {
	service Service
}

func NewTodoRouter(service Service) *TodoRouter {
	return &TodoRouter{service: service}
}

func (todoRouter *TodoRouter) GetHandler() http.Handler {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("GET /todos", todoRouter.getTodos())
	serverMux.HandleFunc("POST /todo/{name}", todoRouter.createTodoRouter())
	serverMux.HandleFunc("DELETE /todo/{id}", todoRouter.deleteTodoRouter())
	return serverMux

}

func (todoRouter *TodoRouter) getTodos() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := todoRouter.service.GetTodos()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dataBytes, err := json.Marshal(todos)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataBytes)

	}
}

func (todoRouter *TodoRouter) createTodoRouter() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.PathValue("name")
		if name := r.PathValue("name"); name != "" {
			todoRouter.service.AddTodo(name)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (todoRouter *TodoRouter) deleteTodoRouter() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if id := r.PathValue("id"); id != "" {
			todoRouter.service.DeleteTodo(id)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
