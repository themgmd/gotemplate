package customerror

type ErrCode int

const (
	UnknownErrorCode = iota
	InvalidCredentialsErrorCode
	InvalidJWTTokenErrorCode
	InvalidOTPCodeErrorCode
	RecordNotFoundErrorCode
)

func (ec ErrCode) IsValid() bool {
	return false
}
