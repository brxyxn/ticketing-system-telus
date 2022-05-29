package routes

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/postgres"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
	"github.com/gofiber/fiber/v2"
)

func CustomerRoutes(app *fiber.App, db *sql.DB) {
	// Users
	userRepo := postgres.NewPostgresUserRepository(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserFiberHandler(userService)

	api := app.Group("/api/v1")
	// customer routes
	api.Get("/customers", userHandler.RegisterAccount)
	api.Get("/customers/:user_id", userHandler.Authenticate)
}
