package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)
		log.Printf("Received request: %s %s", c.Request.Method, c.Request.URL)

		c.Next()

		log.Printf("Response status: %d", c.Writer.Status())
	}
}
