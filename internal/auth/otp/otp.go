package otp

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	DefaultIssuer       = "gotemplate"
	DefaultSecretLength = 25
	DefaultTOTPPeriod   = 30
	DefaultTOPTLength   = 8
	DefaultImageHeight  = 300
	DefaultImageWidth   = 300
)

func Generate(username string) (*otp.Key, error) {
	otpKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      DefaultIssuer,
		AccountName: username,
		Digits:      DefaultTOPTLength,
		Period:      DefaultTOTPPeriod,
		Algorithm:   otp.AlgorithmSHA512,
		SecretSize:  DefaultSecretLength,
	})
	if err != nil {
		return nil, err
	}

	return otpKey, nil
}

func Validate(passcode, secret string) bool {
	return totp.Validate(passcode, secret)
}
