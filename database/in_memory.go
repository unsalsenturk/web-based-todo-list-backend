package database

import (
	"fmt"
	"web-based-todo-list-backend/models"
)

type IDatabase interface {
	GetTodoList() (*models.DataResponse, error)
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

func NewDatabase(todoList models.DataResponse) IDatabase {
	return &Database{todoList: todoList}
}
