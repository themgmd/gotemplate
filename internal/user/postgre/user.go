package postgre

import (
	"context"
	"gotemplate/internal/user/types"
	"gotemplate/pkg/connectors/postgre"
)

type User struct {
	db *postgre.DB
}

func NewUser(db *postgre.DB) *User {
	return &User{db: db}
}

func (u *User) Create(ctx context.Context, user types.User) error {
	_, err := u.db.ExecContext(ctx, createUserQuery, user.Username)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetById() types.User {
	return types.User{}
}

func (u *User) List() string {
	return ""
}

func (u *User) Delete() string {
	return ""
}

func (u *User) Update() string {
	return ""
}
