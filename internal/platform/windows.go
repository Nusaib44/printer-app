package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

func DiscoverPrintersWindows() ([]string, error) {
	cmd := exec.Command("powershell", "Get-Printer")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to discover printers: %v", err)
	}

	// Parse the output to get printer names
	printers := []string{}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "Name") {
			printer := strings.TrimSpace(strings.Split(line, ":")[1])
			printers = append(printers, printer)
		}
	}

	return printers, nil
}
