package platform

import (
	"errors"
	"runtime"
)

// DiscoverPrinters routes to the appropriate OS-specific function.
func DiscoverPrinters() ([]Printer, error) {
	switch runtime.GOOS {
	// case "linux":
	// 	return DiscoverPrintersLinux()
	// case "windows":
	// 	return DiscoverPrintersWindows()
	case "darwin":
		return DiscoverPrintersMacOS()
	default:
		return nil, errors.New("unsupported platform")
	}
}
