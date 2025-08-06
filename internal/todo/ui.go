package todo

type ServiceI interface {
	AddTodo(name string) (id int, err error)
	DeleteTodo(id int) (err error)
	ListTodo() []Todo
	MarkTodo(id int) (err error)
}
