package todo

import (
	"math/rand"
)

type DbHandler interface {
	SaveTodos([]Todo) error
	GetTodos() ([]Todo, error)
}

type Service struct {
	DbHandler DbHandler
}

func (s *Service) AddTodo(name string) error {
	todos, err := s.DbHandler.GetTodos()
	if err != nil {
		return err
	}
	todos = append(todos, Todo{Id: rand.Intn(100), Name: name, Done: false})
	err = s.DbHandler.SaveTodos(todos)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTodo(id int) error {
	todos, err := s.DbHandler.GetTodos()
	if err != nil {
		return err
	}

	filteredTodos := []Todo{}
	for _, v := range todos {
		if v.Id != id {
			filteredTodos = append(filteredTodos, v)
		}
	}

	err = s.DbHandler.SaveTodos(filteredTodos)
	if err != nil {
		return err
	}

	return nil
}
