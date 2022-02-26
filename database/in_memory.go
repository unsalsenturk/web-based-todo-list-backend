package database

import (
	"fmt"
	"web-based-todo-list-backend/models"
)

type IDatabase interface {
	GetTodoList() (*models.DataResponse, error)
	AddTodoList(todo string) (*models.Todo, error)
}

type Database struct {
	todoList models.DataResponse
}

func (db *Database) GetTodoList() (*models.DataResponse, error) {
	if len(db.todoList) == 0 {
		return nil, fmt.Errorf("database Error : db is null")
	}
	return &db.todoList, nil
}
func (db *Database) AddTodoList(todo string) (*models.Todo, error) {
	_, ok := db.todoList[todo]
	if !ok {
		db.todoList[todo] = models.Todo{
			ID:          uint(len(db.todoList)) + 1,
			Description: todo,
		}
		v := db.todoList[todo]
		return &v, nil
	} else {
		return nil, fmt.Errorf("database Error : todo already exist")
	}
}

func NewDatabase(todoList models.DataResponse) IDatabase {
	return &Database{todoList: todoList}
}
