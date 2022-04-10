package routers

import (
	"github.com/gin-gonic/gin"
	todoService "project/src/api/todo/services"
	"project/src/common/middlewares"
)

func MyRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/todos", middlewares.LogMiddleware(), todoService.GetTodos)
	router.GET("/todos/:id", todoService.GetTodo)
	router.POST("/todos", todoService.AddTodo)
	router.PATCH("/todos/:id", todoService.UpdateTodo)
	router.DELETE("/todos/:id", todoService.DeleteTodo)

	router.Run("localhost:9090")
	return router
}
