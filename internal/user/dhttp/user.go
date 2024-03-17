package dhttp

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gotemplate/internal/user/types"
	"gotemplate/pkg/pagination"
)

type UserService interface {
	Create(context.Context, types.User) error
	List(ctx context.Context, pagination pagination.Pagination) ([]types.User, int, error)
}

type User struct {
	user UserService
}

func NewUser(user UserService) *User {
	return &User{user: user}
}

func (u *User) SetupRoutes(router chi.Router) {

}
