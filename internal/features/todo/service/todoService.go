package service

import "cli-todo/internal/features/todo/model"

type Repository interface {
	GetTodo(id string) (model.Todo, error)
	SaveTodo(model.Todo) (model.Todo, error)
	DeleteTodo(model.Todo) error
	UpdateTodo(model.Todo) (model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
}

type TodoService struct {
	repo Repository
}

func NewService(repo Repository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetTodos() ([]model.Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *TodoService) AddTodo(name string) (model.Todo, error) {
	todo, err := s.repo.SaveTodo(model.CreateTodo(name))
	if err != nil {
		return model.Todo{}, err
	}

	return todo, err
}

func (s *TodoService) DeleteTodo(id string) error {
	todo, err := s.repo.GetTodo(id)
	if err != nil {
		return err
	}
	err = s.repo.DeleteTodo(todo)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) ToggleMarkDone(id string) error {
	todo, err := s.repo.GetTodo(id)
	if err != nil {
		return err
	}

	todo.Done = !todo.Done

	_, err = s.repo.UpdateTodo(todo)
	if err != nil {
		return err
	}

	return nil
}
