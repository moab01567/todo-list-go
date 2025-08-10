package todo

import (
	"cli-todo/internal/domainErr"
	"encoding/json"
	"fmt"
	"os"
)

type JsonFileHandler struct {
	filePath string
}

func NewJsonFileHandler(fiePath string) *JsonFileHandler {
	return &JsonFileHandler{filePath: fiePath}
}

func (f *JsonFileHandler) GetTodo(id string) (Todo, error) {
	todos, err := f.readJsonFile()
	if err != nil {
		return Todo{}, domainErr.New("Server Error", "Could not read file", err, domainErr.CodeInternal)
	}

	for _, todo := range todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return Todo{}, domainErr.New(
		"Todo not found",
		fmt.Sprintf("Todo with Id %v not found", id),
		nil,
		domainErr.CodeNotFound)
}

func (f *JsonFileHandler) SaveTodo(saveTodo Todo) (Todo, error) {
	todos, err := f.readJsonFile()
	if err != nil {
		return Todo{}, domainErr.New("Server error", "failed to read file", err, domainErr.CodeInternal)
	}

	todos = append(todos, saveTodo)

	err = f.saveJsonFile(todos)
	if err != nil {
		return Todo{}, domainErr.New("Server Error", "Could not Save file", err, domainErr.CodeInternal)
	}

	return saveTodo, nil
}

func (f *JsonFileHandler) DeleteTodo(deleteTodo Todo) error {
	_, err := f.GetTodo(deleteTodo.Id)
	if err != nil {
		return err
	}
	todos, err := f.readJsonFile()
	if err != nil {
		return domainErr.New("Server error", "failed to read file", err, domainErr.CodeInternal)
	}

	newTodo := []Todo{}
	for _, todo := range todos {
		if todo.Id != deleteTodo.Id {
			newTodo = append(newTodo, todo)
		}
	}

	err = f.saveJsonFile(newTodo)
	if err != nil {
		return domainErr.New("Server Error", "Could not Save file", err, domainErr.CodeInternal)
	}

	return nil

}

func (f *JsonFileHandler) UpdateTodo(updateTodo Todo) (Todo, error) {
	todo, err := f.GetTodo(updateTodo.Id)
	if err != nil {
		return Todo{}, err
	}

	err = f.DeleteTodo(todo)
	if err != nil {
		return Todo{}, err
	}

	updateTodo, err = f.SaveTodo(updateTodo)
	if err != nil {
		return Todo{}, err
	}

	return updateTodo, err
}

func (f *JsonFileHandler) GetAllTodos() ([]Todo, error) {
	todos, err := f.readJsonFile()
	if err != nil {
		return nil, domainErr.New("Server error", "failed to read file", err, domainErr.CodeInternal)
	}

	return todos, nil
}

func (f *JsonFileHandler) readJsonFile() ([]Todo, error) {
	dataBytes, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(dataBytes, &todos)
	if err != nil {
		return nil, err
	}

	return todos, err
}

func (f *JsonFileHandler) saveJsonFile(todos []Todo) error {
	dataBytes, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	err = os.WriteFile(f.filePath, dataBytes, os.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}
	return nil
}
