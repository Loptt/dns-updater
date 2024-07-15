package updater

import "fmt"

// UpdaterDuckDNS is a concrete implementation of the UpdaterInterface object
// that implements the logic to update DDNS records for DuckDNS
// (https://www.duckdns.org/).
type UpdaterDuckDNS struct {
	domain string
}

// Update function performs the action to update the DNS record.
func (u *UpdaterDuckDNS) Update() error {
	fmt.Printf("Updating %s\n", u.domain)

	return nil
}

// MakeUopdaterDuckDNS creates a new UpdaterDuckDNS object.
func MakeUpdaterDuckDNS(domain string) *UpdaterDuckDNS {
	return &UpdaterDuckDNS{
		domain: domain,
	}
}
