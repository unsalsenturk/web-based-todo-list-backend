package service

import (
	"web-based-todo-list-backend/database"
	"web-based-todo-list-backend/models"
)

type IService interface {
	GetTodoList() (*models.ServiceResponse, error)
	AddTodoList(todo string) (*models.Todo, error)
}

type Service struct {
	db database.IDatabase
}

func (svc *Service) GetTodoList() (*models.ServiceResponse, error) {
	res, err := svc.db.GetTodoList()
	if err != nil {
		return nil, err
	}

	srvRes := models.ServiceResponse{}
	for _, todo := range *res {
		srvRes = append(srvRes, todo)
	}
	return &srvRes, nil
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
