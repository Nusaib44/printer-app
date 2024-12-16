package platform

import (
	"bytes"
	"os/exec"
	"strings"
)

type Printer struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Model    string `json:"model"`
}

func DiscoverPrintersMacOS() ([]Printer, error) {
	// Run lpstat to get the list of printers
	cmd := exec.Command("lpstat", "-p")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var printers []Printer
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		// Parse lines starting with "printer"
		if strings.HasPrefix(line, "printer") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				printerName := parts[1]
				// Fetch additional metadata
				status := getPrinterStatus(printerName)
				printerType := detectPrinterType(printerName)
				location := "Local" // Static for now; update if location data is available
				model := getPrinterModel(printerName)

				printers = append(printers, Printer{
					Name:     printerName,
					Status:   status,
					Type:     printerType,
					Location: location,
					Model:    model,
				})
			}
		}
	}

	return printers, nil
}

// Helper function to get the status of a printer
func getPrinterStatus(printerName string) string {
	cmd := exec.Command("lpstat", "-p", printerName)
	output, err := cmd.Output()
	if err != nil {
		return "unknown" // Return "unknown" if unable to fetch status
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

// Helper function to detect printer type (simple logic for demo)
func detectPrinterType(printerName string) string {
	if strings.HasPrefix(printerName, "test") || strings.Contains(printerName, "virtual") {
		return "virtual"
	}
	return "physical"
}

// Helper function to fetch printer model
func getPrinterModel(printerName string) string {
	cmd := exec.Command("lpinfo", "-m", printerName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "unknown"
	}

	// Output looks like "model: HP LaserJet Pro P1102"
	modelInfo := strings.TrimSpace(out.String())
	if modelInfo == "" {
		return "unknown"
	}

	return modelInfo
}
