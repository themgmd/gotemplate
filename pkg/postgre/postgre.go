package postgre

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Options struct {
	DSN         string
	MaxOpenConn int
	MaxIdleConn int
}

type Option func(*Options) error

func DSN(dsn string) Option {
	return func(options *Options) error {
		options.DSN = dsn
		return nil
	}
}

func MaxOpenConn(maxOpenConn int) Option {
	return func(options *Options) error {
		options.MaxOpenConn = maxOpenConn
		return nil
	}
}

func MaxIdleConn(maxIdleConn int) Option {
	return func(options *Options) error {
		options.MaxIdleConn = maxIdleConn
		return nil
	}
}

type DB struct {
	*sqlx.DB
}

func NewConn(ctx context.Context, options ...Option) (*DB, error) {
	var option Options

	for i := range options {
		if err := options[i](&option); err != nil {
			return nil, err
		}
	}

	db, err := sqlx.ConnectContext(ctx, "pgx", option.DSN)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
