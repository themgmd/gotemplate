package redis

import "gotemplate/pkg/redis"

type User struct {
	client *redis.Client
}

func NewUser(client *redis.Client) *User {
	return &User{
		client: client,
	}
}
