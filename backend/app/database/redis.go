package database

import (
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/go-redis/redis"
)

type CacheHandler interface {
	InitializeRedis(bindAddr, password string, dbname int) *redis.Client
}

type CacheHdl struct {
	cache *redis.Client
}

func NewCacheHandler() CacheHandler {
	return &CacheHdl{}
}

func (h *CacheHdl) InitializeRedis(bindAddr, password string, dbname int) *redis.Client {
	h.cache = redis.NewClient(&redis.Options{
		Addr:     bindAddr,
		Password: password,
		DB:       dbname,
	})

	_, err := h.cache.Ping().Result()
	if err != nil {
		h.cache.Close()
		u.Log.Error(err)
		return nil
	}
	u.Log.Info("Connected to cache database", dbname, "with user", password, "at", bindAddr)
	return h.cache
}
