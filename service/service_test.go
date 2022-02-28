package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"web-based-todo-list-backend/mock"
	"web-based-todo-list-backend/models"
)

func TestService_GetTodoList(t *testing.T) {
	t.Run("when GetTodoList returns data properly", func(t *testing.T) {
		databaseReturn := &models.DataResponse{
			"dummy todo": models.Todo{
				ID:          1,
				Description: "buy some milk",
			},
		}
		serviceReturnGetTodoList := &models.ServiceResponse{
			models.Todo{
				ID:          1,
				Description: "buy some milk",
			},
		}
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockDb := mock.NewMockIDatabase(ctl)
		mockDb.
			EXPECT().
			GetTodoList().
			Return(databaseReturn, nil)

		svc := NewService(mockDb)
		res, err := svc.GetTodoList()
		assert.Equal(t, res, serviceReturnGetTodoList)
		assert.Nil(t, err)
	})
	t.Run("when GetTodoList returns error", func(t *testing.T) {
		error := fmt.Errorf("database Error : db is null")
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockDb := mock.NewMockIDatabase(ctl)
		mockDb.
			EXPECT().
			GetTodoList().
			Return(nil, error)

		svc := NewService(mockDb)
		res, err := svc.GetTodoList()
		assert.Equal(t, err, error)
		assert.Nil(t, res)
	})
}
func TestService_AddTodoList(t *testing.T) {
	t.Run("when AddTodoList returns data properly", func(t *testing.T) {
		args := "dummy todo"
		databaseReturn := &models.Todo{
			ID:          1,
			Description: "dummy todo",
		}

		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockDb := mock.NewMockIDatabase(ctl)
		mockDb.
			EXPECT().
			AddTodoList(args).
			Return(databaseReturn, nil)

		svc := NewService(mockDb)
		res, err := svc.AddTodoList(args)
		assert.Equal(t, res, databaseReturn)
		assert.Nil(t, err)
	})
	t.Run("when AddTodoList returns error", func(t *testing.T) {
		args := "dummy todo"
		error := fmt.Errorf("database Error : todo already exist")
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockDb := mock.NewMockIDatabase(ctl)
		mockDb.
			EXPECT().
			AddTodoList(args).
			Return(nil, error)

		svc := NewService(mockDb)
		res, err := svc.AddTodoList(args)
		assert.Equal(t, err, error)
		assert.Nil(t, res)
	})
}
