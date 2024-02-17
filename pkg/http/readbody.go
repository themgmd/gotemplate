package http

import (
	"gotemplate/pkg/customerror"
	"gotemplate/pkg/validator"
	"io"

	"github.com/goccy/go-json"
)

var ErrUnprocessableEntity = customerror.New(customerror.UnprocessableEntityErrorCode, "unprocessable entity")

func ReadBody(body io.ReadCloser, receiver validator.Validator) error {
	err := json.NewDecoder(body).Decode(receiver)
	if err != nil {
		return err
	}

	err = receiver.Validate()
	if err != nil {
		return customerror.Wrap(ErrUnprocessableEntity, err)
	}

	return nil
}
