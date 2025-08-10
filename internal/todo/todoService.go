package todo

type Repository interface {
	GetTodo(id string) (Todo, error)
	SaveTodo(Todo) (Todo, error)
	DeleteTodo(Todo) error
	UpdateTodo(Todo) (Todo, error)
	GetAllTodos() ([]Todo, error)
}

type TodoService struct {
	repo Repository
}

func NewService(repo Repository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetTodos() ([]Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *TodoService) AddTodo(name string) error {
	_, err := s.repo.SaveTodo(CreateTodo(name))
	if err != nil {
		return err
	}
	return nil
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
