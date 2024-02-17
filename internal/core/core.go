package core

import (
	"context"
	"fmt"
	"gotemplate/internal/auth"
	"gotemplate/internal/config"
	"gotemplate/internal/server"
	user "gotemplate/internal/user"
	"gotemplate/migrations"
	"gotemplate/pkg/postgre"
	"net/http"
)

type Core struct {
	httpServer *server.Http
	db         *postgre.DB
}

func New() *Core {
	return &Core{}
}

func (c *Core) Start(ctx context.Context) (err error) {
	cfg := config.Get()

	c.db, err = postgre.New(cfg.Postgre)
	if err != nil {
		return err
	}

	applied, err := migrations.Apply(cfg.Postgre)
	if err != nil {
		return err
	}

	fmt.Printf("Applied migrations: %d\n", applied)

	router := http.NewServeMux()
	auth.Setup(ctx, c.db, router)
	user.Setup(c.db, router)

	c.httpServer = server.NewHttpServer(cfg.HTTP, router)
	go c.httpServer.MustStart()
	return nil
}

func (c *Core) Stop(ctx context.Context) error {
	err := c.httpServer.Stop(ctx)
	if err != nil {
		return err
	}

	err = c.db.Close()
	if err != nil {
		return err
	}

	return nil
}
