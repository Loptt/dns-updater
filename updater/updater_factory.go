package updater

import "fmt"

// UpdaterFactoryInterface provides an abstract interface to create
// UpdaterInterface concrete implementations.
type UpdaterFactoryInterface interface {
	CreateUpdater(string) (UpdaterInterface, error)
}

// UpdaterDuckDNSFactory is a concrete implementation of the
// UpdaterFactoryInterface that creates UpdaterDuckDNS objects.
type UpdaterDuckDNSFactory struct{}

// CreateUpdater creates an updater object of DuckDNS type.
func (uf *UpdaterDuckDNSFactory) CreateUpdater(domain string) (UpdaterInterface, error) {
	if len(domain) < 1 {
		return nil, fmt.Errorf("empty string vwas provided as domain name")
	}
	return MakeUpdaterDuckDNS(domain), nil
}
