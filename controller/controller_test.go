package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"web-based-todo-list-backend/mock"
	"web-based-todo-list-backend/models"
)

func TestController_GetTodoList(t *testing.T) {
	t.Run("when GetTodoList service returns data properly", func(t *testing.T) {
		serviceReturn := &models.DataResponse{
			"dummy todo": models.Todo{
				ID:          1,
				Description: "dummy todo",
			},
		}
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockSvc := mock.NewMockIService(ctl)
		mockSvc.
			EXPECT().
			GetTodoList().
			Return(serviceReturn, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		controller := newTodoListController(mockSvc)
		controller.GetTodoList(c)

		actual := &models.DataResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})
	t.Run("when GetTodoList service returns error", func(t *testing.T) {
		error := fmt.Errorf("database Error : db is null")
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockSvc := mock.NewMockIService(ctl)
		mockSvc.
			EXPECT().
			GetTodoList().
			Return(nil, error)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		controller := newTodoListController(mockSvc)
		controller.GetTodoList(c)

		actual := w.Body.String()

		assert.Equal(t, error.Error(), strings.Trim(actual, "\""))
		assert.Equal(t, w.Result().StatusCode, http.StatusServiceUnavailable)
	})

}
