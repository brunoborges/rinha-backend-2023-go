package cache

import (
	"github.com/filhodanuvem/rinha/internal/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect() error {
	Client = redis.NewClient(&redis.Options{
		Addr: []string{config.RedisURL},
	})

	return nil
}
