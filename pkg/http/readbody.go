package http

import (
	"gotemplate/pkg/customerror"
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
		return customerror.Wrap("unprocessable entity: %w", err)
	}

	return nil
}
