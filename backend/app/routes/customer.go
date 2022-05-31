package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoutes(app *fiber.App, db *sql.DB) {
	// // Users
	// userRepo := postgres.NewPostgresUserRepository(db)
	// userService := customers.NewUserService(userRepo)
	// userHandler := customers.NewUserFiberHandler(userService)

	// api := app.Group("/api/v1")
	// // customer routes
	// api.Get("/customers", userHandler.RegisterAccount)
	// api.Get("/customers/:user_id", userHandler.Authenticate)
}
