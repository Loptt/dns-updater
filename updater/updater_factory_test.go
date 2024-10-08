package updater

import (
	"errors"
	"testing"
)

func TestCreateUpdaterDuckDNS(t *testing.T) {
	tests := []struct {
		description string
		uf          *UpdaterDuckDNSFactory
		domain      string
		token       string
		want_err    error
	}{
		{
			description: "Create with no error",
			uf:          &UpdaterDuckDNSFactory{},
			domain:      "test_domain",
			want_err:    nil,
		},
		{
			description: "Create with empty domain",
			uf:          &UpdaterDuckDNSFactory{},
			domain:      "",
			want_err:    errors.New("domain is empty"),
		},
	}

	for i, test := range tests {
		got, err := test.uf.CreateUpdater(test.domain, test.token)

		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			if got == nil {
				t.Errorf("Test #%d %s: failed to instantiate Updater with error %v", i, test.description, err)
			}
		}
	}
}
