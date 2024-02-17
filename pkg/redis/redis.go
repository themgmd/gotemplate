package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Config interface {
	Addr() string
}

type Client struct {
	*redis.Client
}

func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	ops := &redis.Options{
		Addr: cfg.Addr(),
	}

	client := redis.NewClient(ops)
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
