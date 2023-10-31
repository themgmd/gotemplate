package postgre

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config interface {
	GetDSN() string
	GetMaxIdleConn() int
	GetMaxOpenConn() int
}

type DB struct {
	*sqlx.DB
}

func New(cfg Config) (*DB, error) {
	db, err := sqlx.Open("pgx", cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.GetMaxIdleConn())
	db.SetMaxOpenConns(cfg.GetMaxOpenConn())

	return &DB{db}, nil
}
