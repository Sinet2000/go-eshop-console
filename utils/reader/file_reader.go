package reader

import (
	"fmt"
	"os"
)

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}
