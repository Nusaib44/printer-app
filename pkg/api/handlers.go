package api

import (
	"net/http"
	"printer-app/internal/printers"

	"github.com/gin-gonic/gin"
)

func ListPrintersHandler(c *gin.Context) {
	printers, err := printers.ListPrinters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list printers"})
		return
	}

	c.JSON(http.StatusOK, printers)
}

func PrintDocumentHandler(c *gin.Context) {
	printerName := c.PostForm("printerName")
	filePath := c.PostForm("filePath")

	err := printers.PrintDocument(printerName, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document sent to printer successfully"})
}
