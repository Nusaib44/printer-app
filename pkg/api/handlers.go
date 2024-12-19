package api

import (
	"net/http"
	"printer-app/internal/platform"
	"printer-app/internal/printers"
	"printer-app/pkg/common"

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

func ConfigurePrinter(c *gin.Context) {
	var req struct {
		PrinterName string            `json:"printer_name" binding:"required"`
		Settings    map[string]string `json:"settings"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		common.ErrorResponse(c, err)
		return
	}

	err := platform.Configure(req.PrinterName, req.Settings)
	if err != nil {
		common.ErrorResponse(c, err)
		return
	}

	common.JSONResponse(c, 200, "Printer configuration updated successfully", req)
}
