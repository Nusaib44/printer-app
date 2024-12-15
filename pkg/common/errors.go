package common

import "fmt"

type PrinterError struct {
	Code    int
	Message string
}

func NewPrinterError(code int, message string) error {
	return &PrinterError{Code: code, Message: message}
}

func (e *PrinterError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}
