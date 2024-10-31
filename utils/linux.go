package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDisks() ([]string, error) {
	stdoutStderr, err := exec.Command("df", "-h", "--output=source,target").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error while calling external command: %v", string(stdoutStderr))
	}

	mountPoints := []string{}
	for _, line := range strings.Split(string(stdoutStderr), "\n") {
		if len(line) > 0 && !strings.HasPrefix(line, "Filesystem") {
			parts := strings.Fields(line)
			mountPoints = append(mountPoints, parts[1])
		}
	}

	return mountPoints, nil
}
