package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetVolumes() ([]string, error) {
	cmd := exec.Command("fsutil", "fsinfo", "drives")

	res, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(res))

	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid output format: %s", string(res))
	}
	return fields[1:], nil
}
