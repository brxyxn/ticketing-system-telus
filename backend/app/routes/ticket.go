package routes

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/postgres"
	rd "github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/redis"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/middleware"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/tickets"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func TicketRoutes(app *fiber.App, db *sql.DB, cache *redis.Client) {

	// token initialization
	tokenRepo := rd.NewRedisTokenRepository(cache)         // repository
	tokenService := middleware.NewUserService(tokenRepo)   // service
	middleware := middleware.NewTokenHandler(tokenService) // handler
	// Users
	ticketRepo := postgres.NewPostgresTicketRepository(db)
	ticketService := tickets.NewTicketService(ticketRepo)
	ticketHandler := tickets.NewTicketHandler(ticketService)

	api := app.Group("/api")

	v1 := api.Group("/v1", middleware.Validate)
	v1.Post("/tickets", ticketHandler.CreateTicket)
	v1.Get("/tickets", ticketHandler.GetTickets)
	v1.Get("/tickets/:ticket_id", ticketHandler.GetTicket)
	// v1.Get("/tickets/:ticket_id/comments", ticketHandler.GetTicketComments)

	// // route for tracking of history of tickets
	// api.Get("/tickets/:ticket_id/history", userHandler.RegisterAccount)
}
