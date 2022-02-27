package main

import (
	"log"
	"web-based-todo-list-backend/server"
)

func main() {
	svr := server.NewServer()
	err := svr.StartServer(3000)
	if err != nil {
		log.Fatalln(err)
	}
}
