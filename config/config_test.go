package config

import (
	"errors"
	"testing"

	"github.com/Loptt/dns-updater/file"
)

const sampleGoodConfig = `
intervalseconds: 100
domain: "lopttus"
`

const sampleBadConfig = `
bad config
`

func TestParseConfig(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        Config
		want_err    error
	}{
		{
			description: "Parse normal config",
			input:       sampleGoodConfig,
			want:        Config{IntervalSeconds: 100, Domain: "lopttus"},
			want_err:    nil,
		},
		{
			description: "Parse bad config",
			input:       sampleBadConfig,
			want:        Config{},
			want_err:    errors.New("Error!"),
		},
	}

	for i, test := range tests {
		got, err := parseConfig(test.input)

		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test #%d %s: found error %v", i, test.description, err)
			} else if !got.Equals(test.want) {
				t.Errorf("Test #%d %s: got %v, want %v", i, test.description, got, test.want)
			}
		}
	}
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		description string
		fm          file.FileManagerInterface
		want        Config
		want_err    error
	}{
		{
			description: "Load normal config",
			fm:          &file.FileManagerFake{Content: sampleGoodConfig, Err: nil},
			want:        Config{IntervalSeconds: 100, Domain: "lopttus"},
			want_err:    nil,
		},
		{
			description: "Fail to load config",
			fm:          &file.FileManagerFake{Content: "", Err: errors.New("Error!")},
			want:        Config{IntervalSeconds: 100, Domain: "lopttus"},
			want_err:    errors.New("Error!"),
		},
	}

	for i, test := range tests {
		got, err := LoadConfig("/fake/path", test.fm)

		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			if err != nil {
				t.Errorf("Test #%d %s: found error %v", i, test.description, err)
			} else if !got.Equals(test.want) {
				t.Errorf("Test #%d %s: got %v, want %v", i, test.description, got, test.want)
			}
		}
	}
}
