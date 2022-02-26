package service

import (
	"web-based-todo-list-backend/database"
	"web-based-todo-list-backend/models"
)

type IService interface {
	GetTodoList() (*models.DataResponse, error)
	AddTodoList(todo string) (*models.Todo, error)
}

type Service struct {
	db database.IDatabase
}

func (svc *Service) GetTodoList() (*models.DataResponse, error) {
	res, err := svc.db.GetTodoList()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (svc *Service) AddTodoList(todo string) (*models.Todo, error) {
	res, err := svc.db.AddTodoList(todo)
	if err != nil {
		return nil, err
	}
	return res, err
}

func NewService(database database.IDatabase) IService {
	return &Service{db: database}
}
