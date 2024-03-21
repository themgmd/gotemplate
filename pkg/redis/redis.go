package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Option func(options *redis.Options) error

func Address(addr string) Option {
	return func(options *redis.Options) error {
		options.Addr = addr
		return nil
	}
}

func Password(password string) Option {
	return func(options *redis.Options) error {
		options.Password = password
		return nil
	}
}

func DB(db int) Option {
	return func(options *redis.Options) error {
		options.DB = db
		return nil
	}
}

type Client struct {
	*redis.Client
}

func NewClient(ctx context.Context, options ...Option) (*Client, error) {
	var option redis.Options

	for i := range options {
		if err := options[i](&option); err != nil {
			return nil, err
		}
	}

	client := redis.NewClient(&option)
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
