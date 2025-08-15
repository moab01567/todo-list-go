package repository

import (
	"cli-todo/internal/domainErr"
	"cli-todo/internal/features/todo/model"
	"context"
	"crypto/rand"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqlRepo struct {
	dbFile string
	db     *gorm.DB
}

func NewSqlRepo(dbFile string) (SqlRepo, error) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return SqlRepo{}, err
	}
	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		return SqlRepo{}, err
	}
	return SqlRepo{dbFile: dbFile, db: db}, nil
}

// DeleteTodo implements service.Repository.
func (s SqlRepo) DeleteTodo(model.Todo) error {
	panic("unimplemented")
}

func (s SqlRepo) GetAllTodos() ([]model.Todo, error) {
	ctx := context.Background()
	// id will never stay empty, so id != '' will always return everything
	todos, err := gorm.G[model.Todo](s.db).Where("id != ?", "").Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("filed to get at todos %v", err)
	}
	return todos, err
}

func (s SqlRepo) GetTodo(id string) (model.Todo, error) {
	ctx := context.Background()
	todo, err := gorm.G[model.Todo](s.db).Where("id = ?", id).First(ctx)
	if err != nil {

	}
	return todo, nil
}

func (s SqlRepo) SaveTodo(todo model.Todo) (model.Todo, error) {
	todo.Id = rand.Text()
	ctx := context.Background()
	err := gorm.G[model.Todo](s.db).Create(ctx, &todo)
	if err != nil {
		return model.Todo{}, domainErr.New("Could not create todo", fmt.Sprintf("faild to create todo %#v", todo), err, domainErr.CodeInternal)
	}
	return todo, nil
}

func (s SqlRepo) UpdateTodo(todo model.Todo) (model.Todo, error) {
	s.db.Model(&todo).Updates(todo)
	return todo, nil
}
