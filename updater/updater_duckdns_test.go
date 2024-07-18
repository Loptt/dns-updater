package updater

import (
	"errors"
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
		{
			description: "Update with error in requestor",
			want:        errors.New("Test error"),
			u:           NewUpdaterDuckDNS("test_domain", "test_token", &requestor.RequestorFake{Result: "OK", WantErr: true}),
		},
		{
			description: "Update with wrong return string",
			want:        errors.New("Test error"),
			u:           NewUpdaterDuckDNS("test_domain", "test_token", &requestor.RequestorFake{Result: "BAD", WantErr: false}),
		},
	}

	for i, test := range tests {
		got := test.u.Update()

		if test.want == nil && got != nil {
			t.Errorf("Test #%d %s: want %v got %v", i, test.description, test.want, got)
		} else if test.want != nil && got == nil {
			t.Errorf("Test #%d %s: want %v got %v", i, test.description, test.want, got)
		}
	}
}
