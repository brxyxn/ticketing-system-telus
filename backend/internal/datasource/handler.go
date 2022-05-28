package datsource

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type Handlers struct {
	db    *sql.DB
	cache *redis.Client
}

func NewHandlers(db *sql.DB, cache *redis.Client) *Handlers {
	return &Handlers{db, cache}
}
