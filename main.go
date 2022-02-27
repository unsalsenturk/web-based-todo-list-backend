package main

import (
	"log"
	"web-based-todo-list-backend/controller"
	"web-based-todo-list-backend/database"
	"web-based-todo-list-backend/models"
	"web-based-todo-list-backend/server"
	"web-based-todo-list-backend/service"
)

func main() {
	inMemory := make(models.DataResponse)
	db := database.NewDatabase(inMemory)
	svc := service.NewService(db)
	handler := controller.NewTodoListController(svc)

	svr := server.NewServer()
	err := svr.StartServer(3000, handler)
	if err != nil {
		log.Fatalln(err)
	}
}
