package customerror

type ErrCode int

const (
	UnknownErrorCode = iota
	UnprocessableEntityErrorCode
)

func (ec ErrCode) IsValid() bool {
	return false
}
