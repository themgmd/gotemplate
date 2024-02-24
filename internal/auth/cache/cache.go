package cache

import (
	"context"
	"fmt"
	"gotemplate/internal/auth/types"
	"gotemplate/pkg/cipher"
	"gotemplate/pkg/redis"
	"time"
)

type Cache struct {
	redis *redis.Client
}

func New(redis *redis.Client) *Cache {
	return &Cache{redis: redis}
}

func (c Cache) SaveTempUser(ctx context.Context, user types.CacheUser) (string, error) {
	key, err := cipher.NewKey()
	if err != nil {
		return "", fmt.Errorf("cipher.NewKey: %w", err)
	}

	sessionKey := fmt.Sprintf("%x", key)
	err = c.redis.Set(ctx, sessionKey, &user, time.Minute*15).Err()
	if err != nil {
		return "", fmt.Errorf("c.redis.Set: %w", err)
	}

	return sessionKey, nil
}

func (c Cache) GetTempUser(ctx context.Context, key string) (types.CacheUser, error) {
	var user types.CacheUser
	err := c.redis.Get(ctx, key).Scan(&user)
	if err != nil {
		return user, fmt.Errorf("c.redis.Get: %w", err)
	}

	err = c.redis.Del(ctx, key).Err()
	if err != nil {
		return user, fmt.Errorf("c.redis.Del: %w", err)
	}

	return user, nil
}
