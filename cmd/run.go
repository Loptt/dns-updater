package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Loptt/dns-updater/config"
	"github.com/Loptt/dns-updater/file"
	"github.com/Loptt/dns-updater/scheduler"
	"github.com/Loptt/dns-updater/updater"
)

const defaultConfigPath = "/etc/dns_updater/config.yaml"
const duckDNSEnvVariable = "DUCKDNS_TOKEN"

var configPathFlag = flag.String("config_path", defaultConfigPath, "Path for the configuration file to load")

func Run() error {
	// Parse all flags before using them.
	flag.Parse()

	// Instantiate a real file manager to use to load the config file.
	fm := &file.FileManager{}

	// Load config file that has the values on how to run the DNS updater.
	config, err := config.LoadConfig(*configPathFlag, fm)
	if err != nil {
		return fmt.Errorf("failed to load config with error: %v", err)
	}

	log.Printf("Loaded config from %s, value is %v", *configPathFlag, *config)

	token := os.Getenv(duckDNSEnvVariable)

	uf := &updater.UpdaterDuckDNSFactory{}
	u, err := uf.CreateUpdater(config.Domain, token)
	if err != nil {
		return fmt.Errorf("failed to create updater with error: %v", err)
	}

	s := scheduler.NewScheduler(time.Duration(config.IntervalSeconds)*time.Second, u)
	s.Run()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Run failed with error: %v\n", err)
	}
}
