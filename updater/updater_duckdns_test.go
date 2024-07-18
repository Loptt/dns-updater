package updater

import (
	"testing"

	"github.com/Loptt/dns-updater/requestor"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		description string
		want        error
		u           UpdaterInterface
	}{
		{
			description: "Update with OK status",
			want:        nil,
			u:           NewUpdaterDuckDNS("test_domain", "test_token", &requestor.RequestorFake{Result: "OK", WantErr: false}),
		},
	}

	for i, test := range tests {
		got := test.u.Update()

		if got != test.want {
			t.Errorf("Test #%d %s: want %v got %v", i, test.description, test.want, got)
		}
	}
}
