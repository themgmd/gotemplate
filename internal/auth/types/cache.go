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

func (cu *CacheUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &cu)
}
