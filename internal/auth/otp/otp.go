package otp

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"log/slog"
	"time"
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
	ok, err := totp.ValidateCustom(passcode, secret, time.Now(), totp.ValidateOpts{
		Digits:    DefaultTOPTLength,
		Period:    DefaultTOTPPeriod,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		slog.Error("totp.ValidateCustom", slog.String("error", err.Error()))
	}

	return ok
}
