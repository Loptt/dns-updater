package config

import (
	"fmt"

	"github.com/Loptt/dns-updater/file"
	"gopkg.in/yaml.v3"
)

// Config contains the configuration values for the DNS updater service.
type Config struct {
	IntervalSeconds uint64
	Domain          string
}

func (c Config) Equals(other Config) bool {
	return c.IntervalSeconds == other.IntervalSeconds && c.Domain == other.Domain
}

// parseConfig takes a raw string representing the config in YAML format and
// returns the config struct.
func parseConfig(raw_config string) (*Config, error) {
	config := Config{}
	err := yaml.Unmarshal([]byte(raw_config), &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml content %s, got error: %v", raw_config, err)
	}
	return &config, nil
}

func LoadConfig(path string, fm file.FileManagerInterface) (*Config, error) {
	content, err := fm.Read(path)
	if err != nil {
		return nil, err
	}

	return parseConfig(content)
}
