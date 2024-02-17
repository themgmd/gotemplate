package user

import (
	"context"
	"gotemplate/internal/user/types"
	"gotemplate/pkg/pagination"
)

type Repo interface {
	Create(ctx context.Context, user types.User) error
	GetByLogin(ctx context.Context, login string) (types.User, error)
	List(ctx context.Context, pagination pagination.Pagination) ([]types.User, int, error)
}

type User struct {
	user Repo
}

func New(user Repo) *User {
	return &User{
		user: user,
	}
}

func (u User) Create(ctx context.Context, user types.User) error {
	return u.user.Create(ctx, user)
}

func (u User) GetByLogin(ctx context.Context, login string) (types.User, error) {
	return u.user.GetByLogin(ctx, login)
}

func (u User) List(ctx context.Context, pagination pagination.Pagination) ([]types.User, int, error) {
	return u.user.List(ctx, pagination)
}
