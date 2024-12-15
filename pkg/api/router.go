package api

import (
	"github.com/gin-gonic/gin"
)

// InitializeRouter sets up the HTTP routes for the application
func InitializeRouter(router *gin.Engine) {
	// Define routes
	router.GET("/printers", ListPrintersHandler)
	router.POST("/print", PrintDocumentHandler)
}

// func StartServer(cfg *config.Config) error
func StartServer() error {
	router := gin.Default()
	InitializeRouter(router)
	return router.Run(":" + "3050")
}
