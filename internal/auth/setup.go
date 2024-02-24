package auth

import (
	"context"
	"gotemplate/internal/auth/cache"
	"gotemplate/internal/auth/dhttp"
	"gotemplate/internal/config"
	"gotemplate/internal/user"
	repo "gotemplate/internal/user/repo"
	"gotemplate/pkg/postgre"
	"gotemplate/pkg/redis"
	"log"
	"net/http"
)

func Setup(ctx context.Context, db *postgre.DB, router *http.ServeMux) {
	cfg := config.Get()
	userRepo := repo.NewUser(db)
	client, err := redis.NewClient(ctx, cfg.Redis)
	if err != nil {
		log.Fatalf("redis.NewClient: %s", err.Error())
	}

	authCache := cache.New(client)
	userService := user.New(userRepo)
	service := New(authCache, userService)

	handler := dhttp.NewHandler(service)
	dhttp.NewAuth(handler).SetupRoutes(router)
}
