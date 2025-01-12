package writer

import (
	"fmt"
	"os"
)

// WriteToFile writes data to a file, overwriting any existing content.
func WriteToFile(filePath string, data string) error {
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

// AppendToFile appends data to an existing file.
func AppendToFile(filePath string, data string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for appending: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to append to file: %w", err)
	}

	return nil
}
