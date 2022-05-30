package tickets

type TicketService interface {
	Create(ticket *Ticket) error
	GetTicket(ticketID int64) (Ticket, error)
	GetTickets(userID int64) (Tickets, error)
	CloseTicket(ticket *Ticket) error
}

type ticketService struct {
	database TicketRepository
}

func NewTicketService(database TicketRepository) TicketService {
	return &ticketService{database}
}

func (a *ticketService) Create(ticket *Ticket) error {
	err := a.database.CreateTicket(ticket)
	if err != nil {
		return err
	}
	return nil
}

func (a *ticketService) GetTicket(ticketID int64) (ticket Ticket, err error) {
	ticket.TicketID = ticketID
	err = a.database.GetTicket(&ticket)
	return Ticket{}, err
}

func (a *ticketService) GetTickets(userID int64) (tickets Tickets, err error) {
	err = a.database.GetTickets(&tickets, userID)
	if err != nil {
		return Tickets{}, err
	}
	return nil, err
}

func (a *ticketService) CloseTicket(ticket *Ticket) error {
	err := a.database.CloseTicket(ticket)
	if err != nil {
		return err
	}
	return nil
}
