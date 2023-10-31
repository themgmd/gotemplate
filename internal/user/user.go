package user

import (
	"context"
	"gotemplate/internal/user/types"
	"gotemplate/pkg/pagination"
)

type Repo interface {
	Create(ctx context.Context, user types.User) error
	List(ctx context.Context, pagination pagination.Pagination) ([]types.User, int, error)
}

type Cache interface {
	List(ctx context.Context, key string) ([]types.User, int, error)
}

type User struct {
	user  Repo
	cache Cache
}

func New(user Repo, cache Cache) *User {
	return &User{
		user:  user,
		cache: cache,
	}
}

func (u *User) Create(ctx context.Context, user types.User) error {
	return u.user.Create(ctx, user)
}

func (u *User) List(ctx context.Context, pagination pagination.Pagination) ([]types.User, int, error) {
	// hash pagination
	// get from redis
	// u.cache.List(ctx, hash)

	return u.user.List(ctx, pagination)
}
