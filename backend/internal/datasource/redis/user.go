package redis

import (
	"log"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
	"github.com/go-redis/redis"
)

type userRepository struct {
	cache *redis.Client
}

func NewRedisUserRepository(redis *redis.Client) users.TokenRepository {
	return &userRepository{redis}
}

func (user *userRepository) SetAuthToken(login *users.Login) error {
	err := user.cache.Set(login.Email, login.Token, 0).Err()
	log.Println("setting token:", login.Token)
	return err
}

func (user *userRepository) GetAuthToken(login *users.Login) error {
	err := user.cache.Get(login.Email).Scan(&login.Token)
	log.Println("getting token:", login.Token)
	return err
}
