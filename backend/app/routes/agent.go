package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func AgentRoutes(app *fiber.App, db *sql.DB) {
	// // Users
	// userRepo := postgres.NewPostgresUserRepository(db)
	// userService := customers.NewUserService(userRepo)
	// userHandler := customers.NewUserFiberHandler(userService)

	// api := app.Group("/api/v1")
	// api.Get("/teams/", userHandler.RegisterAccount)
	// api.Get("/teams/:team_id", userHandler.RegisterAccount)
	// api.Get("/tiers", userHandler.RegisterAccount)
	// api.Get("/tiers/:tier_id", userHandler.RegisterAccount)
}
