package todo

import (
	"cli-todo/internal/domainErr"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	serverMux.HandleFunc("POST /todo/n/{name}", todoRouter.createTodoRouter())
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
		name := strings.Trim("", r.PathValue("name"))

		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := todoRouter.service.AddTodo(name); err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (todoRouter *TodoRouter) deleteTodoRouter() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if id := r.PathValue("id"); id != "" {
			err := todoRouter.service.DeleteTodo(id)
			w.WriteHeader(domainErr.GetHttpStatus(err))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
