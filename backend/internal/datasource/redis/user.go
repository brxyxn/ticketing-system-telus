package redis

import (
	"time"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
	"github.com/go-redis/redis"
)

// TODO: request device information like browser, os, device, etc. to be stored in redis
// instead of the IP address.

type tokenRepository struct {
	cache *redis.Client
}

func NewRedisUserRepository(redis *redis.Client) users.TokenRepository {
	return &tokenRepository{redis}
}

func (t *tokenRepository) SetAuthToken(login *users.Login) error {
	err := t.cache.Set(login.Email+login.IP, login.Token, 2*time.Minute).Err()
	return err
}
