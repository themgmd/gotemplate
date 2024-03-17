package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"gotemplate/internal/auth"
	"gotemplate/internal/config"
	"gotemplate/internal/middleware"
	"gotemplate/internal/user"
	"gotemplate/migrations"
	"gotemplate/pkg/errors"
	"gotemplate/pkg/healthcheck"
	"gotemplate/pkg/postgre"
	"gotemplate/pkg/redis"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func runApp() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("config initialized")
	conf := config.Get()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	slog.Info("applying migrations")
	applied, err := migrations.Apply(conf.Postgre.DSN())
	if err != nil {
		return err
	}

	slog.Info("migrations successfully applied", "applied", applied)

	slog.Info("initialize postgres connection")
	dbConn, err := postgre.NewConn(
		ctx,
		postgre.DSN(conf.Postgre.DSN()),
		postgre.MaxOpenConn(conf.Postgre.MaxOpenConn),
		postgre.MaxIdleConn(conf.Postgre.MaxIdleConn),
	)
	if err != nil {
		return err
	}

	defer func() { err = errors.JoinCloser(err, dbConn) }()

	redisClient, err := redis.NewClient(
		ctx,
		redis.Address(conf.Redis.Host),
		redis.Password(conf.Redis.Password),
	)
	if err != nil {
		return err
	}

	defer func() { err = errors.JoinCloser(err, redisClient) }()

	slog.Info("initialize router")
	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	slog.Info("apply middleware")
	middleware.NewManager(router).ApplyMiddleware()

	slog.Info("initialize auth module")
	auth.Setup(ctx, dbConn, redisClient, apiRouter)

	slog.Info("initialize user module")
	user.Setup(dbConn, apiRouter)

	router.Mount("/api", apiRouter)

	slog.Info("server initialized")
	server := &http.Server{
		Addr:              conf.HTTP.Host,
		Handler:           router,
		ReadHeaderTimeout: conf.HTTP.ReadHeaderTimeout,
	}

	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(2)

	eg.Go(func() error {
		defer slog.Info("stop listen http server")

		slog.Info("init http server", "host", conf.HTTP.Host)
		healthcheck.Get().MarkAsUp()
		serr := server.ListenAndServe()
		if errors.Is(serr, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("group listed http server: %w", serr)
	})

	eg.Go(func() error {
		defer slog.Info("shutdown http server")

		<-ctx.Done()
		healthcheck.Get().MarkAsDown()
		serr := server.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("group shutdown http server: %w", serr)
		}
		return nil
	})

	if err = eg.Wait(); err != nil {
		return err
	}

	slog.Info("shutdown app")
	time.Sleep(time.Second * 1)

	return nil
}
