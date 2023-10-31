package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

func Open(ctx context.Context) (*Client, error) {
	ops := &redis.Options{}

	client := redis.NewClient(ops)
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
