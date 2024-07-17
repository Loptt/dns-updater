package requestor

import (
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRequestorHTTPRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description string
		r           RequestorInterface
		want        string
		want_err    error
		setup       func()
		cleanup     func()
	}{
		{
			description: "Request with 200 status",
			r:           &RequestorHttp{url: "http://test.com"},
			want:        "OK",
			want_err:    nil,
			setup: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", "http://test.com", httpmock.NewStringResponder(200, "OK"))
			},
			cleanup: func() { httpmock.DeactivateAndReset() },
		},
		{
			description: "Request with 200 status and empty body",
			r:           &RequestorHttp{url: "http://test.com"},
			want:        "",
			want_err:    nil,
			setup: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", "http://test.com", httpmock.NewStringResponder(200, ""))
			},
			cleanup: func() { httpmock.DeactivateAndReset() },
		},
		{
			description: "Request with non-OK status returns error",
			r:           &RequestorHttp{url: "http://test.com"},
			want:        "",
			want_err:    errors.New("request failed with non-ok status"),
			setup: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", "http://test.com", httpmock.NewStringResponder(401, ""))
			},
			cleanup: func() { httpmock.DeactivateAndReset() },
		},
	}

	for i, test := range tests {
		test.setup()
		got, err := test.r.Request()
		test.cleanup()

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
