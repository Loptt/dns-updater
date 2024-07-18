package updater

import (
	"fmt"

	"github.com/Loptt/dns-updater/requestor"
)

// UpdaterDuckDNSFactory is a factory object that creates UpdaterDuckDNS
// objects.
type UpdaterDuckDNSFactory struct{}

// CreateUpdater creates an updater object of DuckDNS type.
func (uf *UpdaterDuckDNSFactory) CreateUpdater(domain, token string) (UpdaterInterface, error) {
	if len(domain) < 1 {
		return nil, fmt.Errorf("empty string vwas provided as domain name")
	}
	return NewUpdaterDuckDNS(domain, token, &requestor.RequestorHttp{}), nil
}
