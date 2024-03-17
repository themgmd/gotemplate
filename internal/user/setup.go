package user

import (
	"github.com/go-chi/chi/v5"
	"gotemplate/internal/user/dhttp"
	userPostgre "gotemplate/internal/user/repo"
	"gotemplate/pkg/postgre"
)

func Setup(db *postgre.DB, router chi.Router) *User {
	userRepo := userPostgre.NewUser(db)
	userService := New(userRepo)
	userHandler := dhttp.NewUser(userService)
	userHandler.SetupRoutes(router)

	return userService
}
