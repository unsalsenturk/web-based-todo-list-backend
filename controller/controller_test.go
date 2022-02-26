package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
	"web-based-todo-list-backend/mock"
)

func TestController_GetTodoList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockSvc := mock.NewMockIService(ctl)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// need controller

}
