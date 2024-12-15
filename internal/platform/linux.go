package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

func DiscoverPrintersLinux() ([]string, error) {
	cmd := exec.Command("lpstat", "-e")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to discover printers: %v", err)
	}

	// Parse the output to get printer names
	printers := []string{}
	for _, line := range strings.Split(string(out), "\n") {
		if line != "" {
			printer := strings.Fields(line)[0]
			printers = append(printers, printer)
		}
	}

	return printers, nil
}
