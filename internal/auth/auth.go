package auth

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gotemplate/internal/auth/hasher"
	"gotemplate/internal/auth/otp"
	"gotemplate/internal/auth/types"
	"gotemplate/internal/config"
	userTypes "gotemplate/internal/user/types"
	"gotemplate/pkg/cipher"
	"image/png"
)

type Cache interface {
	SaveTempUser(ctx context.Context, user types.CacheUser) (string, error)
	GetTempUser(ctx context.Context, key string) (types.CacheUser, error)
}

type User interface {
	Create(ctx context.Context, user userTypes.User) error
	GetByLogin(ctx context.Context, login string) (userTypes.User, error)
}

type Auth struct {
	cache  Cache
	user   User
	cipher *cipher.Cipher
}

func New(cache Cache, user User) *Auth {
	return &Auth{
		cache:  cache,
		user:   user,
		cipher: cipher.New([]byte(config.Get().App.EncryptionKey)),
	}
}

func (a Auth) InitRegistration(
	ctx context.Context,
	request types.InitRegistrationRequest,
) (resp types.InitRegistrationResponse, err error) {
	cacheUser := types.CacheUser{
		Username: request.Username,
		Email:    request.Email,
	}

	cacheUser.Password, err = hasher.GenerateFromPassword(request.Password)
	if err != nil {
		err = fmt.Errorf("hasher.GenerateFromPassword: %w", err)
		return
	}

	otpKey, err := otp.Generate(cacheUser.Username)
	if err != nil {
		err = fmt.Errorf("otp.Generate: %w", err)
		return
	}

	resp.Secret = otpKey.Secret()
	cacheUser.OTPSecret, err = a.cipher.Encode(otpKey.Secret())
	if err != nil {
		err = fmt.Errorf("a.cipher.Encode: %w", err)
		return
	}

	resp.Identifier, err = a.cache.SaveTempUser(ctx, cacheUser)
	if err != nil {
		err = fmt.Errorf("a.cache.SaveTempUser: %w", err)
		return
	}

	qrCode, err := otpKey.Image(otp.DefaultImageWidth, otp.DefaultImageHeight)
	if err != nil {
		err = fmt.Errorf("otpKey.Image: %w", err)
		return
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode)
	if err != nil {
		err = fmt.Errorf("png.Encode: %w", err)
		return
	}

	resp.QrCode = buf.Bytes()
	return
}

func (a Auth) FinishRegistration(
	ctx context.Context,
	request types.FinishRegistrationRequest,
) (err error) {
	cacheUser, err := a.cache.GetTempUser(ctx, request.Identifier)
	if err != nil {
		err = fmt.Errorf("a.cache.GetTempUser: %w", err)
		return
	}

	otpSecret, err := a.cipher.Decode(cacheUser.OTPSecret)
	if err != nil {
		err = fmt.Errorf("a.cipher.Decode: %w", err)
		return
	}

	if !otp.Validate(request.Code, otpSecret) {
		err = errors.New("invalid otp code")
		return
	}

	newUser := userTypes.NewUser(cacheUser.Username, cacheUser.Email)
	newUser.Password = cacheUser.Password
	newUser.OTPSecret = cacheUser.OTPSecret

	err = a.user.Create(ctx, *newUser)
	if err != nil {
		err = fmt.Errorf("a.user.Create: %w", err)
		return
	}

	return
}

func (a Auth) Login(ctx context.Context, request types.LoginRequest) error {
	user, err := a.user.GetByLogin(ctx, request.Login)
	if err != nil {
		return fmt.Errorf("a.user.GetByLogin: %w", err)
	}

	err = hasher.CompareWithPassword(request.Password, user.Password)
	if err != nil {
		return fmt.Errorf("hasher.CompareWithPassword: %w", err)
	}

	otpSecret, err := a.cipher.Decode(user.OTPSecret)
	if err != nil {
		return fmt.Errorf("a.cipher.Decode: %w", err)
	}

	if !otp.Validate(request.Code, otpSecret) {
		return errors.New("invalid otp code")
	}

	// todo: jwt token generate

	return nil
}
