package gophertunneloracle

import (
	"bytes"
	"fmt"
	"os/exec"
)

func gitCommand(directory string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	if directory != "" {
		cmd.Dir = directory
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		message := stderr.String()
		if message == "" {
			message = stdout.String()
		}
		return "", fmt.Errorf("git %v: %w: %s", args, err, message)
	}
	return stdout.String(), nil
}
