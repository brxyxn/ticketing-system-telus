package tickets

type TicketRepository interface {
	CreateTicket(ticket *Ticket) error               // create new ticket
	GetTicket(ticket *Ticket) error                  // get ticket by id
	GetTickets(tickets *Tickets, userID int64) error // list all tickets
	CloseTicket(ticket *Ticket) error                // update ticket status to closed
}
