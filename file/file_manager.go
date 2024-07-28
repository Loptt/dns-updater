package file

import (
	"fmt"
	"os"
)

// FileManager object provices concrete implementations to perform file actions
// (read, write, etc.)
type FileManager struct{}

// Read function reads contents from the file specified in `path` and returns
// it as a string.
func (fm *FileManager) Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s with error: %v", path, err)
	}

	return string(data), nil
}
