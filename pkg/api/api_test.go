package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"printer-app/internal/config"
	"printer-app/internal/printers"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListPrintersHandler(t *testing.T) {
	cfg := &config.Config{ServerPort: "8080"}
	log.Print(cfg)
	r := gin.Default()
	printers.RegisterRoutes(r)

	req, _ := http.NewRequest("GET", "/printers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", w.Code)
	}

	var response []printers.Printer
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response: %v", err)
	}

	if len(response) == 0 {
		t.Error("Expected some printers in response, got none")
	}
}

func TestPrintDocumentHandler(t *testing.T) {
	cfg := &config.Config{ServerPort: "8080"}
	log.Print(cfg)
	r := gin.Default()
	printers.RegisterRoutes(r)

	form := url.Values{}
	form.Add("printerName", "Printer1")
	form.Add("filePath", "/path/to/test/file")

	req, _ := http.NewRequest("POST", "/print", bytes.NewBufferString(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response: %v", err)
	}

	expectedMessage := "Document sent to printer successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, response["message"])
	}
}
