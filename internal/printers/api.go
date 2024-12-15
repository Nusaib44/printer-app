package printers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/printers", ListPrintersHandler)
	router.POST("/print", PrintDocumentHandler)
}

func ListPrintersHandler(c *gin.Context) {
	printers, err := ListPrinters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, printers)
}

func PrintDocumentHandler(c *gin.Context) {
	printerName := c.PostForm("printerName")
	filePath := c.PostForm("filePath")

	err := PrintDocument(printerName, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document sent to printer successfully"})
}
