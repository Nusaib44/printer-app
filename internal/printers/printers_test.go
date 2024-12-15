package printers

import (
	"testing"
)

func TestListPrinters(t *testing.T) {
	printers, err := ListPrinters()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(printers) == 0 {
		t.Error("Expected some printers, got none")
	}
}

func TestPrintDocument(t *testing.T) {
	err := PrintDocument("Printer1", "/path/to/test/file")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
