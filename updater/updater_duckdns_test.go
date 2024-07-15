package updater

import "testing"

func TestUpdate(t *testing.T) {
	tests := []struct {
		description string
		want        error
		u           UpdaterInterface
	}{
		{
			description: "Update with no error",
			want:        nil,
			u:           MakeUpdaterDuckDNS("test_domain"),
		},
	}

	for i, test := range tests {
		got := test.u.Update()

		if got != test.want {
			t.Errorf("Test #%d %s: want %v got %v", i, test.description, test.want, got)
		}
	}
}
