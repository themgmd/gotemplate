package dhttp

import (
	"context"
	"gotemplate/internal/user/types"
	"gotemplate/pkg/transport"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService interface {
	Create(context.Context, types.User) error
}

type User struct {
	user UserService
}

func NewUser(user UserService) *User {
	return &User{user: user}
}

func (u *User) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/user", u.Create).Methods(http.MethodPost)
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var user types.User

	err := transport.ReadBody(r.Body, &user)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	err = u.user.Create(r.Context(), user)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(transport.NewSuccessResponse(user).Bytes())
}
