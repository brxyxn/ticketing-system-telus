package tickets

type TicketService interface {
}

type ticketService struct {
	database TicketRepository
}
