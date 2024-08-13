package secret

import (
	"fmt"
	"os"
)

// SecretManagerEnv implements the SecretManagerInterface to fetch secrets
// from the environment variables.
type SecretManagerEnv struct{}

// Fetch takes the key as the environment variable and fetches its value.
func (sm *SecretManagerEnv) Fetch(key string) (string, error) {
	result, found := os.LookupEnv(key)
	if !found {
		return "", fmt.Errorf("key %s was not found in env", key)
	}

	return result, nil
}

func NewSecretManagerEnv() *SecretManagerEnv {
	return &SecretManagerEnv{}
}
