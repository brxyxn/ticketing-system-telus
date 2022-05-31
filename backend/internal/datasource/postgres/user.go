package postgres

import (
	"database/sql"

	"github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/customers"
)

type userRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) customers.UserRepository {
	return &userRepository{db}
}

const (
	qGetCompanyID         = `SELECT company_id FROM companies WHERE name = $1`
	qInsertCompany        = `INSERT INTO companies (name) VALUES ($1) RETURNING company_id`
	qInsertCustomer       = `INSERT INTO customers (company_id) VALUES ($1) RETURNING customer_id`
	qDeleteCustomerOnFail = `DELETE FROM customers WHERE customer_id = $1`
	qInsertUser           = `INSERT INTO users (email, password, customer_id) VALUES ($1, $2, $3) RETURNING user_id`
	qInsertProfile        = `INSERT INTO profiles (first_name, last_name, user_id) VALUES ($1, $2, $3) RETURNING *`
)

func (user *userRepository) CreateAccount(account *customers.Account) error {
	var err error
	// validate if company exists, if not create it
	account.Customer.CompanyID = 0
	err = user.db.QueryRow(
		qGetCompanyID, &account.Company.Name,
	).Scan(&account.Company.CompanyID)
	if err != sql.ErrNoRows && err != nil {
		return err
	}

	if account.Company.CompanyID == 0 {
		err = user.db.QueryRow(qInsertCompany, &account.Company.Name).Scan(&account.Company.CompanyID)
	}
	account.Customer.CompanyID = account.Company.CompanyID

	// TODO: validate if profile exists and has user_id assigned and email address is the same as the input, if not create it or reject registration
	// create customer and profile
	err = user.db.QueryRow(
		qInsertCustomer,
		account.Company.CompanyID,
	).Scan(&account.Customer.CustomerID)
	if err != nil {
		return err
	}

	account.User.CustomerId = account.Customer.CustomerID

	// create user
	err = user.db.QueryRow(
		qInsertUser,
		account.User.Email, account.User.Password, account.User.CustomerId,
	).Scan(&account.User.UserID)
	if err != nil {
		utils.Log.Debug(&account.User.Email, &account.User.Password, &account.User.CustomerId)
		user.db.QueryRow(qDeleteCustomerOnFail, account.Customer.CustomerID)
		return err
	}

	// create profile
	err = user.db.QueryRow(
		qInsertProfile,
		account.Profile.FirstName, account.Profile.LastName, account.User.UserID,
	).Scan(&account.Profile.ProfileID, &account.Profile.FirstName, &account.Profile.LastName, &account.Profile.UserID, &account.Profile.CreatedAt, &account.Profile.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (user *userRepository) GetAccountByEmail(email string) (*customers.Account, error) {
	var account customers.Account

	err := user.db.QueryRow(
		`select p.profile_id,
			p.first_name,
			p.last_name,
			u.user_id,
			u.email,
			u.password,
			u.status,
			cx.customer_id,
			c.name
		from profiles p
			inner join users u on u.user_id = p.user_id
			inner join customers cx on cx.customer_id = u.customer_id
			inner join companies c on c.company_id = cx.company_id
		where u.email = $1`, email,
	).Scan(
		&account.Profile.ProfileID,
		&account.Profile.FirstName,
		&account.Profile.LastName,
		&account.User.UserID,
		&account.User.Email,
		&account.User.Password,
		&account.User.Status,
		&account.Customer.CustomerID,
		&account.Company.Name,
	)
	if err != nil {
		return &customers.Account{}, err
	}

	return &account, nil
}
