package postgres

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/tickets"
)

type ticketRepository struct {
	db *sql.DB
}

func NewPostgresTicketRepository(db *sql.DB) tickets.TicketRepository {
	return &ticketRepository{db}
}

const (
	qCreateTicket  = `INSERT INTO tickets (title, description, case_id, status_id) VALUES ($1, $2, $3, $4) RETURNING ticket_id`
	qGetTicketByID = `SELECT * FROM tickets WHERE ticket_id = $1`
	qGetTickets    = `SELECT * FROM tickets`
	qCloseTicket   = `UPDATE tickets SET status_id = 3 WHERE ticket_id = $1`
)

func (t *ticketRepository) CreateTicket(ticket *tickets.Ticket) error {
	err := t.db.QueryRow(
		qCreateTicket, ticket.Title, ticket.Description, ticket.CaseID, ticket.StatusID,
	).Scan(
		&ticket.TicketID,
	)
	return err
}

func (t *ticketRepository) GetTicket(ticket *tickets.Ticket) error {
	err := t.db.QueryRow(
		qGetTicketByID,
		ticket.TicketID,
	).Scan(
		&ticket.TicketID,
		&ticket.Title,
		&ticket.Description,
		&ticket.StatusID,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)
	return err
}

// TODO: add pagination

func (t *ticketRepository) GetTickets(ts *tickets.Tickets, userID int64) error {
	rows, err := t.db.Query(qGetTickets)
	if err != nil {
		return err
	}

	ts2 := make(tickets.Tickets, 0)
	var ticket tickets.Ticket
	for rows.Next() {
		if err := rows.Scan(
			ticket.TicketID,
			ticket.Title,
			ticket.Description,
			ticket.CaseID,
			ticket.StatusID,
			ticket.CreatedAt,
			ticket.UpdatedAt,
			ticket.ClosedAt,
		); err != nil {
			return err
		}
		ts2 = append(ts2, ticket)
	}

	ts = &ts2

	defer rows.Close()
	return nil
}

func (t *ticketRepository) CloseTicket(ticket *tickets.Ticket) error {
	_, err := t.db.Exec(
		qCloseTicket,
		ticket.TicketID,
	)
	return err
}
