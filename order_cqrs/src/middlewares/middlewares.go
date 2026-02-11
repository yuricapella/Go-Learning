package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Logger - escreve as informações no terminal
func Logger(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		fmt.Printf("\n %s %s %s\n", ginContext.Request.Method, ginContext.Request.RequestURI, ginContext.Request.Host)
		next(ginContext)
	}
}
