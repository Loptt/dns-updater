package updater

import (
	"fmt"

	"github.com/Loptt/dns-updater/requestor"
)

// UpdaterDuckDNS is a concrete implementation of the UpdaterInterface object
// that implements the logic to update DDNS records for DuckDNS
// (https://www.duckdns.org/).
type UpdaterDuckDNS struct {
	domain string
	token  string
	r      requestor.RequestorInterface
}

// DuckDNS URL to update the DNS record.
const duckDnsUpdateUrl = "https://www.duckdns.org/update"

// DuckDNS string indicating a successful request.
const duckDnsValidResponse = "OK"

// Update function performs the action to update the DNS record.
func (u *UpdaterDuckDNS) Update() error {
	fullUrl := fmt.Sprintf("%s?domains=%s&token=%s", duckDnsUpdateUrl, u.domain, u.token)
	result, err := u.r.Request(fullUrl)
	if err != nil {
		return fmt.Errorf("failed to request to domain %s: %v", u.domain, err)
	}

	if result != duckDnsValidResponse {
		return fmt.Errorf("response returned by DuckDNS is invalid for domain %s, got %s, expected %s", u.domain, result, duckDnsValidResponse)
	}

	return nil
}

// NewUpdaterDuckDNS creates a new UpdaterDuckDNS object.
func NewUpdaterDuckDNS(domain, token string, r requestor.RequestorInterface) *UpdaterDuckDNS {
	return &UpdaterDuckDNS{
		domain: domain,
		token:  token,
		r:      r,
	}
}
