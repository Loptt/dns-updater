package config

// Config contains the configuration values for the DNS updater service.
type Config struct {
	intervalSeconds uint64
	domain          string
}

func (c Config) Equals(other Config) bool {
	return c.intervalSeconds == other.intervalSeconds && c.domain == other.domain
}

// ParseConfig takes a raw string representing the config in YAML format and
// returns the config struct.
func ParseConfig(raw_config string) Config {
	//TODO(Loptt): Remove hardcoded values and parse them from the file. This
	// needs the file package to be ready.
	return Config{intervalSeconds: 86400, domain: "lopttus"}
}
