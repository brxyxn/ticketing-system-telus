package tickets

type Tickets []Ticket

type Ticket struct {
	TicketID    int64  `json:"ticket_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CaseID      int64  `json:"case_id"`
	StatusID    int64  `json:"status_id"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	ClosedAt    string `json:"closed_at"`
}

type Status struct {
	StatusID  int64  `json:"status_id"`
	Title     string `json:"title"`
	ByDefault bool   `json:"by_default"`
}

type Comments []Comment

type Comment struct {
	CommentID int64  `json:"comment_id"`
	Body      string `json:"body"`
	TicketID  int64  `json:"ticket_id"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"-"`
}

type Cases []Case

type Case struct {
	CaseID   int64 `json:"case_id"`
	Assigned bool  `json:"assigned"`
	Status   bool  `json:"status"`
	AgentID  int64 `json:"agent_id"`
	TierID   int64 `json:"tier_id"`
}

type Incident struct {
	UserID int64 `json:"user_id"`
	CaseID int64 `json:"case_id"`
}

type UserTickets struct {
	Tickets   []Ticket   `json:"tickets"`
	Cases     []Case     `json:"cases"`
	Incidents []Incident `json:"incidents"`
}
