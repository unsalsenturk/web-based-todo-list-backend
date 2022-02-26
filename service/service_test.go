package service

import (
	"github.com/golang/mock/gomock"
	"testing"
	"web-based-todo-list-backend/mock"
	"web-based-todo-list-backend/models"
)

func TestService_GetTodoList(t *testing.T) {
	databaseReturn := &models.DataResponse{
		"dummy todo": models.Todo{
			ID:          1,
			Description: "dummy todo",
		},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockDb := mock.NewMockIDatabase(ctl)
	mockDb.
		EXPECT().
		GetTodoList().
		Return(databaseReturn, nil)

	// need service
}
