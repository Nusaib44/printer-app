package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

func DiscoverPrintersWindows() ([]Printer, error) {
	cmd := exec.Command("wmic", "printer", "get", "Name,Status,Location,DriverName")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var printers []Printer
	lines := strings.Split(string(output), "\n")
	for i, line := range lines {
		// Skip header
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			printerName := parts[0]
			status := parts[1]
			model := "unknown" // Placeholder, can enhance if more data is available

			printers = append(printers, Printer{
				Name:     printerName,
				Status:   status,
				Type:     "physical",
				Location: "unknown", // Windows doesn't always provide this
				Model:    model,
			})
		}
	}

	return printers, nil
}

func configureForWindows(printerName string, settings map[string]string) error {
	// Example: Set default printer using PowerShell
	if settings["default"] == "true" {
		cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`Set-DefaultPrinter -Name "%s"`, printerName))
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to set default printer: %w", err)
		}
	}
	// Add more settings logic as needed
	return nil
}
