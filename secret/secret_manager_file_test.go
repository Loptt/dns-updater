package secret

import (
	"errors"
	"testing"

	"github.com/Loptt/dns-updater/file"
)

func TestFetch(t *testing.T) {
	tests := []struct {
		description string
		sm          SecretManagerInterface
		fm          file.FileManagerInterface
		want        string
		want_err    error
	}{
		{
			description: "Fetch existing key",
			sm:          &SecretManagerFile{f: &file.FileManagerFake{Content: "secret", Err: nil}},
			want:        "secret",
			want_err:    nil,
		},
		{
			description: "Fail to read file",
			sm:          &SecretManagerFile{f: &file.FileManagerFake{Content: "", Err: errors.New("Error!")}},
			want:        "",
			want_err:    errors.New("Error!"),
		},
	}

	for i, test := range tests {
		got, err := test.sm.Fetch("/fake/path")

		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test #%d %s: found error %v", i, test.description, err)
			} else if got != test.want {
				t.Errorf("Test #%d %s: got %v, want %v", i, test.description, got, test.want)
			}
		}
	}
}
