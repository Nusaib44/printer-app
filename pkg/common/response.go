package common

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, err error) {
	code := 500
	if e, ok := err.(*PrinterError); ok {
		code = e.Code
	}
	JSONResponse(c, code, err.Error(), nil)
}
