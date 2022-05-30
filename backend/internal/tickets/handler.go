package tickets

import (
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/gofiber/fiber/v2"
)

type TicketHandler interface {
	CreateTicket(c *fiber.Ctx) error
	GetTicket(c *fiber.Ctx) error
	GetTickets(c *fiber.Ctx) error
	CloseTicket(c *fiber.Ctx) error
}

type ticketHandler struct {
	service TicketService
}

func NewTicketHandler(service TicketService) TicketHandler {
	return &ticketHandler{service}
}

func (a *ticketHandler) CreateTicket(c *fiber.Ctx) error {
	u.Log.Info("Creating new ticket")
	var ticket Ticket
	err := c.BodyParser(&ticket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if ticket.Title == "" || ticket.Description == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing credentials"})
	}

	err = a.service.Create(&ticket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(ticket)
}

func (a *ticketHandler) GetTicket(c *fiber.Ctx) error {
	u.Log.Info("Getting ticket by id")
	var ticketID int64
	err := c.BodyParser(&ticketID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing ticket id"})
	}

	ticket, err := a.service.GetTicket(ticketID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(ticket)
}

func (a *ticketHandler) GetTickets(c *fiber.Ctx) error {
	u.Log.Info("Getting all tickets")
	var userID int64
	err := c.BodyParser(&userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing user id"})
	}

	tickets, err := a.service.GetTickets(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tickets)
}

func (a *ticketHandler) CloseTicket(c *fiber.Ctx) error {
	u.Log.Info("Closing ticket")
	var ticket Ticket
	err := c.BodyParser(&ticket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing ticket id"})
	}

	err = a.service.CloseTicket(&ticket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "ticket closed"})
}
