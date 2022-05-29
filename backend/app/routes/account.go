package routes

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/postgres"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api/v1")

	userRepo := postgres.NewPostgresUserRepository(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserFiberHandler(userService)

	// Accounts
	api.Post("/register", userHandler.RegisterAccount) // register profile, user and [customer + company | agent + team]
	api.Post("/login", userHandler.RegisterAccount)    // login user
	api.Get("/user", userHandler.RegisterAccount)      // returns user profile
	api.Get("/logout", userHandler.RegisterAccount)    // logout user
}
