package tickets

type TicketRepository interface {
	GetTickets() (Tickets, error)               // list all tickets
	GetTicket(id int64) (Ticket, error)         // get ticket by id
	CreateTicket(ticket Ticket) (Ticket, error) // create new ticket
	CloseTicket(ticket Ticket) (Ticket, error)  // update ticket status to closed
}