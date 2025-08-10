package todo

import (
	"cli-todo/internal/appError"
	"cli-todo/internal/domainErr"
	"encoding/json"
	"os"
)

type JsonFileHandler struct {
	filePath string
}

func NewJsonFileHandler(fiePath string) *JsonFileHandler {
	return &JsonFileHandler{filePath: fiePath}
}
func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, appError.New("could not read file", err)
	}

	return data, err
}
func (f *JsonFileHandler) DeleteTodoFromDB(id string) error {

	return nil

}

func (f *JsonFileHandler) GetTodos() ([]Todo, error) {
	data, err := readFile(f.filePath)
	if err != nil {
		return nil, domainErr.New(string(domainErr.CodeInternal), err, domainErr.CodeInternal)
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, domainErr.New(string(domainErr.CodeInternal), err, domainErr.CodeInternal)
	}
	return todos, nil
}

func (f *JsonFileHandler) SaveTodos(todos []Todo) error {
	var data []byte
	var err error
	data, err = json.Marshal(todos)
	if err != nil {
		return appError.New("Could not pars struct to byte", err)
	}

	err = os.WriteFile(f.filePath, data, os.FileMode(os.O_CREATE))
	if err != nil {
		return appError.New("Could not write to file", err)
	}
	return nil
}
