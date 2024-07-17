package requestor

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// RequestorHttp implements the concrete RequestorInterface that sends HTTP
// requests to remote services to perform the Request() action.
type RequestorHttp struct {
	url string
}

// Encodes the status of a successful HTTP request.
const statusOk = 200

// Calls an HTTP GET to the `url` and returns the body of the request as a
// string.
func (r *RequestorHttp) Request() (string, error) {
	resp, err := http.Get(r.url)
	if err != nil {
		return "", fmt.Errorf("failed to do GET request at url: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != statusOk {
		return "", fmt.Errorf("request returned status code %d", resp.StatusCode)
	}

	buffer := &strings.Builder{}

	n, err := io.Copy(buffer, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse request body: %v", err)
	}

	log.Printf("Read %d bytes from request body\n", n)

	return buffer.String(), nil
}
