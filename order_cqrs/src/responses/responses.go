package responses

import (
	"net/http"

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

// Created returns a 201 Created response with Location header and ID
func Created(ginContext *gin.Context, location string, id interface{}) {
	ginContext.Header("Location", location)
	ginContext.JSON(http.StatusCreated, gin.H{"id": id})
}
