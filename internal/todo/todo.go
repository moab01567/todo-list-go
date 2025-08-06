package todo

type Todo struct {
	Id   int
	Name string
	Done bool
}

func (s *Todo) ToggleDone() {
	s.Done = !s.Done
}
