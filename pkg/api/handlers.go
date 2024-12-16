package api

import (
	"net/http"
	"printer-app/internal/platform"
	"printer-app/internal/printers"

	"github.com/gin-gonic/gin"
)

func ListPrintersHandler(c *gin.Context) {
	printers, err := platform.DiscoverPrinters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"printers": printers})
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
