package main

import (
	"fmt"
	"time"

	"github.com/Loptt/dns-updater/scheduler"
	"github.com/Loptt/dns-updater/updater"
)

func main() {
	uf := &updater.UpdaterDuckDNSFactory{}

	domain := "lopttus"
	token := "test_token"
	interval := 1 * time.Second
	u, err := uf.CreateUpdater(domain, token)
	if err != nil {
		fmt.Printf("Failed to create updater with error: %v\n", err)
	}

	s := scheduler.NewScheduler(interval, u)

	s.Run()
}
