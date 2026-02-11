package responses

import (
	"github.com/gin-gonic/gin"
)

// JSON returns a JSON response for the request
func JSON(ginContext *gin.Context, statusCode int, data interface{}) {
	ginContext.JSON(statusCode, data)
}

// Error returns an error in JSON format for the request
func Error(ginContext *gin.Context, statusCode int, err error) {
	ginContext.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
