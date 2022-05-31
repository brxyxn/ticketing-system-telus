package redis

import (
	"github.com/brxyxn/ticketing-system-telus/backend/internal/middleware"
	"github.com/go-redis/redis"
)

// middleware
func NewRedisTokenRepository(redis *redis.Client) middleware.TokenRepository {
	return &tokenRepository{redis}
}

func (t *tokenRepository) GetAuthToken(login *middleware.Login) error {
	err := t.cache.Get(login.Email + login.IP).Scan(&login.Token)
	return err
}
