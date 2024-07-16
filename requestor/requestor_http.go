package requestor

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type RequestorHttp struct {
	url string
}

const statusOk = 200

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
