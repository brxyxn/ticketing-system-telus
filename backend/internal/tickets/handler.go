package tickets

type TicketHandler interface {
	// CreateTicket(c *fiber.Ctx) error
}

type ticketHandler struct {
	service TicketService
}

func NewTicketHandler(service TicketService) TicketHandler {
	return &ticketHandler{service}
}
