package todo

import (
	"cli-todo/appError"
	"encoding/json"
	"os"
)

type JsonFileHandler struct {
	FilePath string
}

func (f *JsonFileHandler) GetTodos() ([]Todo, error) {
	var data []byte
	var todos []Todo
	var err error

	data, err = os.ReadFile(f.FilePath)
	if err != nil {
		return nil, appError.New("could not read file", err)
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, appError.New("Could not pars json", err)
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

	err = os.WriteFile(f.FilePath, data, os.FileMode(os.O_CREATE))
	if err != nil {
		return appError.New("Could not write to file", err)
	}
	return nil
}
