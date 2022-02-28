package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"web-based-todo-list-backend/mock"
	"web-based-todo-list-backend/models"
)

func TestController_GetTodoList(t *testing.T) {
	t.Run("when GetTodoList service returns data properly", func(t *testing.T) {
		serviceReturn := &models.ServiceResponse{
			models.Todo{
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

		controller := NewTodoListController(mockSvc)
		controller.GetTodoList(c)

		actual := &models.ServiceResponse{}
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

		controller := NewTodoListController(mockSvc)
		controller.GetTodoList(c)

		actual := w.Body.String()

		assert.Equal(t, error.Error(), strings.Trim(actual, "\""))
		assert.Equal(t, w.Result().StatusCode, http.StatusServiceUnavailable)
	})

}
func TestController_AddTodoList(t *testing.T) {
	t.Run("when AddTodoList service returns data properly", func(t *testing.T) {
		args := "dummy todo"
		serviceReturn := &models.Todo{
			ID:          1,
			Description: "dummy todo",
		}
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockSvc := mock.NewMockIService(ctl)
		mockSvc.
			EXPECT().
			AddTodoList(args).
			Return(serviceReturn, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := &http.Request{
			URL:    &url.URL{},
			Header: make(http.Header), // if you need to test headers
		}

		r := io.NopCloser(strings.NewReader("{\"todo\":\"dummy todo\"}"))
		req.Body = r
		c.Request = req

		controller := NewTodoListController(mockSvc)
		controller.AddTodoList(c)

		actual := &models.Todo{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
	})
	t.Run("when AddTodoList service returns error", func(t *testing.T) {
		args := "dummy todo"
		error := fmt.Errorf("database Error : todo already exist")
		ctl := gomock.NewController(t)
		defer ctl.Finish()
		mockSvc := mock.NewMockIService(ctl)
		mockSvc.
			EXPECT().
			AddTodoList(args).
			Return(nil, error)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := &http.Request{
			URL:    &url.URL{},
			Header: make(http.Header), // if you need to test headers
		}

		r := io.NopCloser(strings.NewReader("{\"todo\":\"dummy todo\"}"))
		req.Body = r
		c.Request = req

		controller := NewTodoListController(mockSvc)
		controller.AddTodoList(c)

		actual := w.Body.String()

		assert.Equal(t, error.Error(), strings.Trim(actual, "\""))
		assert.Equal(t, w.Result().StatusCode, http.StatusServiceUnavailable)
	})

}
