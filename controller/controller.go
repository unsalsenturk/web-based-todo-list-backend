package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-based-todo-list-backend/models"
	"web-based-todo-list-backend/service"
)

type IController interface {
	GetTodoList(c *gin.Context)
	AddTodoList(c *gin.Context)
}

type Controller struct {
	svc service.IService
}

func (ctl *Controller) GetTodoList(c *gin.Context) {
	res, err := ctl.svc.GetTodoList()
	if err != nil {
		switch err.Error() {
		case "database Error : db is null":
			c.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func (ctl *Controller) AddTodoList(c *gin.Context) {
	body := models.PostTodoListBody{}
	if err := c.BindJSON(&body); err != nil {
		return
	}

	res, err := ctl.svc.AddTodoList(body.Todo)
	if err != nil {
		switch err.Error() {
		case "database Error : todo already exist":
			c.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func newTodoListController(service service.IService) IController {
	return &Controller{svc: service}
}
