package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Middleware called 2")
		context.Next()
	}
}
