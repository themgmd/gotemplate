package http

import (
	"io"
	"net/http"
)

// Do makes openapi requests at the given url with
// specific method, headers and request body
func Do(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	cl := http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
