package platform

import (
	"bytes"
	"os/exec"
	"strings"
)

// DiscoverPrintersLinux lists all available printers using the lpstat command.
func DiscoverPrintersLinux() ([]string, error) {
	cmd := exec.Command("lpstat", "-p")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var printers []string
	for _, line := range lines {
		if strings.HasPrefix(line, "printer ") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				printers = append(printers, fields[1])
			}
		}
	}
	return printers, nil
}
