package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func TicketRoutes(app *fiber.App, db *sql.DB) {
	// // Users
	// userRepo := postgres.NewPostgresUserRepository(db)
	// userService := customers.NewUserService(userRepo)
	// userHandler := customers.NewUserFiberHandler(userService)

	// api := app.Group("/api/v1")
	// api.Get("/tickets", userHandler.RegisterAccount)
	// api.Get("/tickets/:ticket_id", userHandler.RegisterAccount)
	// api.Get("/tickets/:ticket_id/comments", userHandler.RegisterAccount)

	// // route for tracking of history of tickets
	// api.Get("/tickets/:ticket_id/history", userHandler.RegisterAccount)
}
