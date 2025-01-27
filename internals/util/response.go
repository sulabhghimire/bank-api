package util

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func SuccessResponse(msg string, data any) gin.H {
	return gin.H{"message": msg, "data": data}
}
