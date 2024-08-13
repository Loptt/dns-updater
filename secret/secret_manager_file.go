package secret

import (
	"fmt"

	"github.com/Loptt/dns-updater/file"
)

// SecretManagerFile implements the SecretManagerInterface to fetch secrets
// from the file system.
type SecretManagerFile struct {
	f file.FileManagerInterface
}

// Fetch takes the key as the file path, and returns its contents.
func (sm *SecretManagerFile) Fetch(key string) (string, error) {
	result, err := sm.f.Read(key)
	if err != nil {
		return "", fmt.Errorf("failed to fetch secret from file %s with error: %v", key, err)
	}

	return result, nil
}

func NewSecretManagerFile(f file.FileManagerInterface) *SecretManagerFile {
	return &SecretManagerFile{f: f}
}
