package repo

import (
	"context"
	"fmt"
	"gotemplate/internal/user/types"
	"gotemplate/pkg/pagination"
	"gotemplate/pkg/postgre"
)

type User struct {
	db *postgre.DB
}

func NewUser(db *postgre.DB) *User {
	return &User{db: db}
}

func (u *User) Create(ctx context.Context, user types.User) error {
	_, err := u.db.ExecContext(
		ctx,
		createUserQuery,
		user.ID,
		user.Email,
		user.Username,
		user.OTPSecret,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetByLogin(ctx context.Context, login string) (types.User, error) {
	var user types.User

	err := u.db.GetContext(ctx, &user, getByLogin, login)
	if err != nil {
		return user, fmt.Errorf("u.db.GetContext: %w", err)
	}

	return user, nil
}

func (u *User) List(ctx context.Context, pagination pagination.Pagination) ([]types.User, int, error) {
	var (
		total int
		users []types.User
	)

	err := u.db.SelectContext(ctx, users, getUserListQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		return users, total, err
	}

	return users, total, nil
}
