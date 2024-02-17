package types

import "gotemplate/pkg/postgre"

type User struct {
	*postgre.BaseModel
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	OTPSecret string `json:"-" db:"otp_secret"`
}

func NewUser(username, email string) *User {
	return &User{
		BaseModel: postgre.NewBaseModel(),
		Username:  username,
		Email:     email,
	}
}
