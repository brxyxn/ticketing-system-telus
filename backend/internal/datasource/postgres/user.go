package postgres

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
)

type userRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) users.UserRepository {
	return &userRepository{db}
}

func (user *userRepository) CreateAccount(account *users.Account) error {
	// base
	user.db.QueryRow("INSERT INTO users (email, password, first_name, last_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", account.Email, account.Password, account.FirstName, account.LastName)

	return nil
}

func (user *userRepository) GetAccountByEmail(email string) (*users.Account, error) {
	// _, err := user.db.Exec()

	return nil, nil
}
