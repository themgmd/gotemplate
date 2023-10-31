package user

import (
	"github.com/gorilla/mux"
	"gotemplate/internal/user/dhttp"
	userPostgre "gotemplate/internal/user/postgre"
	"gotemplate/pkg/postgre"
)

func Setup(db *postgre.DB, router *mux.Router) *User {
	userRepo := userPostgre.NewUser(db)
	userService := New(userRepo)
	userHandler := dhttp.NewUser(userService)
	userHandler.SetupRoutes(router)

	return userService
}
