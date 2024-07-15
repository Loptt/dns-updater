package updater

import "fmt"

type UpdaterDuckDNS struct {
	domain string
}

func (u *UpdaterDuckDNS) Update() error {
	fmt.Printf("Updating %s\n", u.domain)

	return nil
}

func MakeUpdaterDuckDNS(domain string) *UpdaterDuckDNS {
	return &UpdaterDuckDNS{
		domain: domain,
	}
}
