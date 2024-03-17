package auth

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gotemplate/internal/auth/cache"
	"gotemplate/internal/auth/dhttp"
	"gotemplate/internal/user"
	"gotemplate/internal/user/repo"
	"gotemplate/pkg/postgre"
	"gotemplate/pkg/redis"
)

func Setup(
	_ context.Context,
	db *postgre.DB,
	client *redis.Client,
	router chi.Router,
) {
	userRepo := repo.NewUser(db)

	authCache := cache.New(client)
	userService := user.New(userRepo)
	service := New(authCache, userService)

	handler := dhttp.NewHandler(service)
	dhttp.NewAuth(handler).SetupRoutes(router)
}
