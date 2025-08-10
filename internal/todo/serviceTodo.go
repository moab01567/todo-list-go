package todo

type Repository interface {
	GetTodo(id string) (Todo, error)
	SaveTodo(Todo) (Todo, error)
	DeleteTodo(Todo) error
	UpdateTodo(Todo) (Todo, error)
	GetAllTodos() ([]Todo, error)
}

type ServiceTodo struct {
	repo Repository
}

func NewService(repo Repository) *ServiceTodo {
	return &ServiceTodo{repo: repo}
}

func (s *ServiceTodo) GetTodos() ([]Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *ServiceTodo) AddTodo(name string) error {
	_, err := s.repo.SaveTodo(CreateTodo(name))
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceTodo) DeleteTodo(id string) error {
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
