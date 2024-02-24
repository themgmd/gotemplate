package types

import (
	"github.com/goccy/go-json"
)

type CacheUser struct {
	Email     string
	Username  string
	Password  string
	OTPSecret string
}

func (cu *CacheUser) MarshalBinary() (data []byte, err error) {
	return json.Marshal(cu)
}

func (cu *CacheUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &cu)
}
