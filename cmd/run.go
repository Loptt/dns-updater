package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Loptt/dns-updater/config"
	"github.com/Loptt/dns-updater/file"
	"github.com/Loptt/dns-updater/scheduler"
	"github.com/Loptt/dns-updater/secret"
	"github.com/Loptt/dns-updater/updater"
)

const defaultConfigPath = "/etc/dns_updater/config.yaml"
const duckDNSSecretFile = "/run/secrets/duckdns_token"

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

	// SecretManagerInterface is used to fetch the token used to authenticate
	// with the DDNS service.
	sm := secret.NewSecretManagerFile(fm)
	token, err := sm.Fetch(duckDNSSecretFile)
	if err != nil {
		return fmt.Errorf("failed to get secret token with error %v", err)
	}

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
