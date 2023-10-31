package dhttp

import (
	"context"
	"gotemplate/internal/user/types"
	httplib "gotemplate/pkg/http"
	"gotemplate/pkg/pagination"
	"net/http"

	"github.com/gorilla/mux"
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

func (u *User) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/user", u.create).Methods(http.MethodPost)
	router.HandleFunc("/user", u.list).Methods(http.MethodGet)
}

func (u *User) create(w http.ResponseWriter, r *http.Request) {
	var user types.User

	err := httplib.ReadBody(r.Body, &user)
	if err != nil {
		httplib.NewErrorResponse(w, httplib.StatusBadRequest, err)
		return
	}

	err = u.user.Create(r.Context(), user)
	if err != nil {
		httplib.NewErrorResponse(w, httplib.StatusInternalServerError, err)
		return
	}

	httplib.NewSuccessResponse(w, httplib.StatusCreated, user)
}

func (u *User) list(w http.ResponseWriter, r *http.Request) {
	var pag pagination.RequestPagination

	err := httplib.ReadQuery(r, &pag)
	if err != nil {
		httplib.NewErrorResponse(w, httplib.StatusBadRequest, err)
		return
	}

	users, total, err := u.user.List(r.Context(), *pag.ToPagination())
	if err != nil {
		httplib.NewErrorResponse(w, httplib.StatusInternalServerError, err)
		return
	}

	meta := *pagination.NewResponsePagination(pag, total)
	httplib.NewListResponse(w, meta, users)
}
