package handler

import (
	"cli-todo/internal/domainErr"
	"cli-todo/internal/features/todo/model"
	"cli-todo/internal/httpserver"
	"net/http"
)

type Service interface {
	AddTodo(name string) (model.Todo, error)
	DeleteTodo(id string) error
	GetTodos() ([]model.Todo, error)
	ToggleMarkDone(id string) error
}

type TodoRouter struct {
	service Service
}

func NewTodoHandler(service Service) *TodoRouter {
	return &TodoRouter{service: service}
}

func (todoRouter *TodoRouter) GetHandler() http.Handler {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("GET /todos", todoRouter.getTodos())
	serverMux.HandleFunc("POST /todo/n/{name}", todoRouter.createTodoRouter())
	serverMux.HandleFunc("DELETE /todo/{id}", todoRouter.deleteTodoRouter())
	serverMux.HandleFunc("PUT /todo/done/mark/{id}", todoRouter.toggleMarkDone())
	return serverMux

}

func (todoRouter *TodoRouter) getTodos() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := todoRouter.service.GetTodos()

		if err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
			return
		}

		dataBytes, err := httpserver.TypeToBytes(todos)
		if err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataBytes)

	}
}

func (todoRouter *TodoRouter) createTodoRouter() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, err := todoRouter.service.AddTodo(name)
		if err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
			return
		}

		dataBytes, err := httpserver.TypeToBytes(todo)
		if err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataBytes)
	}
}

func (todoRouter *TodoRouter) deleteTodoRouter() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id != "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := todoRouter.service.DeleteTodo(id)
		if err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (todoRouter *TodoRouter) toggleMarkDone() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := todoRouter.service.ToggleMarkDone(id)
		if err != nil {
			w.WriteHeader(domainErr.GetHttpStatus(err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}

}
