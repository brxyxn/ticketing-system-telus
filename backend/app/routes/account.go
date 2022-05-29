package routes

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/customers"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/postgres"
	rd "github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/redis"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/middleware"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(app *fiber.App, db *sql.DB, cache *redis.Client) {

	// token initialization
	tokenRepo := rd.NewRedisTokenRepository(cache)         // repository
	tokenService := middleware.NewUserService(tokenRepo)   // service
	middleware := middleware.NewTokenHandler(tokenService) // handler
	// user initialization
	postgresRepo := postgres.NewPostgresUserRepository(db)            // repositories
	uRedisRepo := rd.NewRedisUserRepository(cache)                    // repositories
	userService := customers.NewUserService(postgresRepo, uRedisRepo) // services
	userHandler := customers.NewUserHandler(userService)              // controllers/handlers

	api := app.Group("/api")

	// Accounts for customers
	api.Post("/customer/register", userHandler.RegisterAccount) // register profile, user and [customer + company | agent + team]
	api.Post("/customer/login", userHandler.Login)              // login user

	// v1 routes
	v1 := api.Group("/v1", middleware.Validate)

	v1.Get("/customer/user", userHandler.GetUser) // returns user profile

	// Accounts for agents
	v1.Post("/agent/register", userHandler.RegisterAccount) // register profile, user and [customer + company | agent + team]
	v1.Post("/agent/login", userHandler.Login)              // login user
	v1.Get("/agent/user", userHandler.RegisterAccount)      // returns user profile

	v1.Get("/logout", userHandler.RegisterAccount) // logout user
}
