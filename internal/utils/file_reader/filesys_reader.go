package file_reader

import (
	"fmt"
	"os"
)

// FileSystemReader implements the FileReader interface for file system operations.
type FileSystemReader struct{}

func (fsr *FileSystemReader) ReadFile(filePath string) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file from filesystem: %w", err)
	}
	return fileContent, nil
}