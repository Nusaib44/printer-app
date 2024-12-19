package platform

import (
	"fmt"
	"runtime"
)

type Printer struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Model    string `json:"model"`
}

// DiscoverPrinters detects printers based on the OS
func DiscoverPrinters() ([]Printer, error) {
	switch runtime.GOOS {
	case "darwin":
		return DiscoverPrintersMacOS()
	case "linux":
		return DiscoverPrintersLinux()
	case "windows":
		return DiscoverPrintersWindows()
	default:
		return nil, nil // Unsupported OS
	}
}

func Configure(printerName string, settings map[string]string) error {
	switch runtime.GOOS {
	case "darwin", "linux":
		return configureForUnix(printerName, settings)
	case "windows":
		return configureForWindows(printerName, settings)
	default:
		return fmt.Errorf("unsupported platform")
	}
}
