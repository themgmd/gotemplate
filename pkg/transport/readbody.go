package transport

import (
	"io"

	"github.com/goccy/go-json"
)

func ReadBody(body io.ReadCloser, receiver interface{}) error {
	err := json.NewDecoder(body).Decode(receiver)
	if err != nil {
		return err
	}

	// TODO: validation
	return nil
}
