package user

import (
	"context"
	"gotemplate/internal/user/types"
)

type UserRepo interface {
	Create(ctx context.Context, user types.User) error
}

type User struct {
	user UserRepo
}

func New(user UserRepo) *User {
	return &User{user: user}
}

func (u *User) Create(ctx context.Context, user types.User) error {
	return u.user.Create(ctx, user)
}
