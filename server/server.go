package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-based-todo-list-backend/controller"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int, handler controller.IController) error {

	gin.SetMode(gin.ReleaseMode)
	var router = gin.New()
	router.Use(corsMiddleware())

	v1 := router.Group("api/v1")
	{
		v1.GET("todolist", handler.GetTodoList)
		v1.POST("todolist", handler.AddTodoList)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	err := srv.ListenAndServe()
	return err

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://35.222.60.250:80/")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			// 204 No Content
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
