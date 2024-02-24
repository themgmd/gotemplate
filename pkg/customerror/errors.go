package customerror

var (
	ErrUserNotExist       = New(RecordNotFoundErrorCode, "user not found")
	ErrInvalidCredentials = New(InvalidCredentialsErrorCode, "invalid credentials")
)
