package routes

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/postgres"
	rd "github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/redis"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(app *fiber.App, db *sql.DB, cache *redis.Client) {
	api := app.Group("/api/v1")

	// repositories
	postgresRepo := postgres.NewPostgresUserRepository(db)
	redisRepo := rd.NewRedisUserRepository(cache)
	// services
	userService := users.NewUserService(postgresRepo, redisRepo)
	// controllers/handlers
	userHandler := users.NewUserFiberHandler(userService)

	// Accounts for customers
	api.Post("/customer/register", userHandler.RegisterAccount) // register profile, user and [customer + company | agent + team]
	api.Post("/customer/login", userHandler.Authenticate)       // login user
	api.Get("/customer/user", userHandler.GetUser)              // returns user profile

	// Accounts for agents
	api.Post("/agent/register", userHandler.RegisterAccount) // register profile, user and [customer + company | agent + team]
	api.Post("/agent/login", userHandler.Authenticate)       // login user
	api.Get("/agent/user", userHandler.RegisterAccount)      // returns user profile

	api.Get("/logout", userHandler.RegisterAccount) // logout user
}
