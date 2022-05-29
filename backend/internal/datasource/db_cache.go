package datsource

import (
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/go-redis/redis"
)

func (h *Handlers) InitializeCache(bindAddr, password string, dbname int) {
	h.cache = redis.NewClient(&redis.Options{
		Addr:     bindAddr,
		Password: password,
		DB:       dbname,
	})

	_, err := h.cache.Ping().Result()
	if err != nil {
		h.cache.Close()
		u.Log.Error(err)
		return
	}
	u.Log.Info("Connected to cache database", dbname, "with user", password, "at", bindAddr)
}
