package todo

import (
	"crypto/rand"
)

type Todo struct {
	Id   string
	Name string
	Done bool
}

func CreateRandomTodo(name string) Todo {
	return Todo{Id: rand.Text(),
		Name: name,
		Done: false}
}

func (s *Todo) ToggleDone() {
	s.Done = !s.Done
}
