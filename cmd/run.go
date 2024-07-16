package main

import (
	"fmt"

	"github.com/Loptt/dns-updater/updater"
)

func main() {
	uf := &updater.UpdaterDuckDNSFactory{}

	domain := "lopttus"
	u, err := uf.CreateUpdater(domain)
	if err != nil {
		fmt.Printf("Failed to create updater with error: %v\n", err)
	}

	if err := u.Update(); err != nil {
		fmt.Printf("Failed to update with error: %v\n", err)
	}

	fmt.Printf("Update successful\n")
}
