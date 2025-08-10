package todo

type DbHandler interface {
	SaveTodos([]Todo) error
	GetTodos() ([]Todo, error)
}

type Service struct {
	dbHandler DbHandler
}

func NewService(dbHandler DbHandler) *Service {
	return &Service{dbHandler: dbHandler}
}

func (s *Service) GetTodos() ([]Todo, error) {
	return s.dbHandler.GetTodos()
}

func (s *Service) AddTodo(name string) error {
	todos, err := s.dbHandler.GetTodos()
	if err != nil {
		return err
	}
	todos = append(todos, CreateTodo(name))
	err = s.dbHandler.SaveTodos(todos)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTodo(id string) error {
	todos, err := s.dbHandler.GetTodos()
	if err != nil {
		return err
	}

	filteredTodos := []Todo{}
	for _, v := range todos {
		if v.Id != id {
			filteredTodos = append(filteredTodos, v)
		}
	}

	err = s.dbHandler.SaveTodos(filteredTodos)
	if err != nil {
		return err
	}

	return nil
}
