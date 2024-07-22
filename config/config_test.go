package config

import "testing"

func TestParseConfig(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        Config
	}{
		{
			description: "Parse normal config",
			input:       "test",
			want:        Config{intervalSeconds: 86400, domain: "lopttus"},
		},
	}

	for i, test := range tests {
		got := ParseConfig(test.input)

		if !got.Equals(test.want) {
			t.Errorf("Test #%d %s: want %v got %v", i, test.description, test.want, got)
		}
	}
}
