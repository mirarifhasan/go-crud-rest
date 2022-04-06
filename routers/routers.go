package routers

import (
	"github.com/gin-gonic/gin"
	todoService "project/src/todo/services"
)

func MyRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/todos", todoService.GetTodos)
	router.GET("/todos/:id", todoService.GetTodo)
	router.POST("/todos", todoService.AddTodo)
	router.PATCH("/todos/:id", todoService.UpdateTodo)
	router.DELETE("/todos/:id", todoService.DeleteTodo)

	router.Run("localhost:9090")
	return router
}
