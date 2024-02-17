package user

import (
	"gotemplate/internal/user/dhttp"
	userPostgre "gotemplate/internal/user/repo"
	"gotemplate/pkg/postgre"
	"net/http"
)

func Setup(db *postgre.DB, router *http.ServeMux) *User {
	userRepo := userPostgre.NewUser(db)
	userService := New(userRepo)
	userHandler := dhttp.NewUser(userService)
	userHandler.SetupRoutes(router)

	return userService
}
