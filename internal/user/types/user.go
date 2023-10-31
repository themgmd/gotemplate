package types

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID         uuid.UUID `json:"id" db:"id""`
	FamilyName string    `json:"familyName" db:"family_name"`
	GivenName  string    `json:"givenName" db:"given_name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	RegDate    time.Time `json:"regDate" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt  null.Time `json:"deletedAt" db:"deleted_at"`
}

func NewUser(familyName, givenName, email string) *User {
	createdAt := time.Now()
	return &User{
		ID:         uuid.New(),
		FamilyName: familyName,
		GivenName:  givenName,
		Email:      email,
		RegDate:    createdAt,
		UpdatedAt:  createdAt,
	}
}

func (u *User) SetPassword(passwd string) error {
	hPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hPasswd)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
