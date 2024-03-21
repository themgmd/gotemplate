package http

import (
	"gotemplate/pkg/errors"
	"gotemplate/pkg/validator"
	"io"

	"github.com/goccy/go-json"
)

func ReadBody(body io.ReadCloser, receiver validator.Validator) error {
	err := json.NewDecoder(body).Decode(receiver)
	if err != nil {
		return err
	}

	err = receiver.Validate()
	if err != nil {
		return errors.Wrap(err, "validation error")
	}

	return nil
}
