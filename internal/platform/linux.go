package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

func DiscoverPrintersLinux() ([]Printer, error) {
	cmd := exec.Command("lpstat", "-p")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var printers []Printer
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "printer") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				printerName := parts[1]
				status := getPrinterStatusLinux(printerName)
				model := getPrinterModelLinux(printerName)

				printers = append(printers, Printer{
					Name:     printerName,
					Status:   status,
					Type:     "physical", // Assuming physical for simplicity
					Location: "Local",    // Can add logic to fetch location
					Model:    model,
				})
			}
		}
	}

	return printers, nil
}

func getPrinterStatusLinux(printerName string) string {
	cmd := exec.Command("lpstat", "-p", printerName)
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	if strings.Contains(string(output), "idle") {
		return "idle"
	}
	if strings.Contains(string(output), "printing") {
		return "printing"
	}
	if strings.Contains(string(output), "offline") {
		return "offline"
	}
	return "unknown"
}

func getPrinterModelLinux(printerName string) string {
	cmd := exec.Command("lpinfo", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, printerName) {
			return line
		}
	}
	return "unknown"
}

func configureForUnix(printerName string, settings map[string]string) error {
	// Example: Set default printer
	if settings["default"] == "true" {
		cmd := exec.Command("lpoptions", "-d", printerName)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to set default printer: %w", err)
		}
	}
	// Add more settings logic as needed
	return nil
}
