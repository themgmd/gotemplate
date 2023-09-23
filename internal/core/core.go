package core

import (
	"context"
	"gotemplate/internal/config"
	"gotemplate/internal/server"
	user "gotemplate/internal/user"
	"gotemplate/pkg/connectors/postgre"

	"github.com/gorilla/mux"
)

type Core struct {
	httpServer *server.Http
	db         *postgre.DB
}

func New() *Core {
	return &Core{}
}

func (c *Core) Run(_ context.Context) (err error) {
	cfg := config.Get()

	c.db, err = postgre.New(cfg.Postgre)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
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
