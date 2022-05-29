package users

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Account struct {
	User
	Profile
	Customer
	Company
}

type User struct {
	UserID     int64  `json:"user_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Status     bool   `json:"status"`
	CustomerId int64  `json:"customer_id"`
	AgentId    int64  `json:"agent_id"`
	Admin      bool   `json:"admin"`
	CreatedAt  string `json:"-"`
	UpdatedAt  string `json:"-"`
}

type Profile struct {
	ProfileID int64  `json:"profile_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

type Customer struct {
	CustomerID int64 `json:"customer_id"`
	CompanyID  int64 `json:"company_id"`
}

type Company struct {
	CompanyID int64  `json:"company_id"`
	Name      string `json:"name"`
}
