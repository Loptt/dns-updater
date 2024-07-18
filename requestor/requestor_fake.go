package requestor

import "fmt"

// RequestorFake provides a fake implementation of a RequestorInterface
// object.
type RequestorFake struct {
	Result  string
	WantErr bool
}

// Request returns the specified string by r.Result. If r.WantErr is true, then
// the function returns an error.
func (r *RequestorFake) Request(url string) (string, error) {
	if r.WantErr {
		return "", fmt.Errorf("request error")
	}
	return r.Result, nil
}
