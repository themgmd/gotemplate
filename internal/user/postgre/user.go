package postgre

import (
	"context"
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
		user.FamilyName,
		user.GivenName,
		user.Email,
		user.Password,
		user.RegDate,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetById() types.User {
	return types.User{}
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

func (u *User) Delete() string {
	return ""
}

func (u *User) Update() string {
	return ""
}
