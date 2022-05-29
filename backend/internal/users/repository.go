package users

// postgres
type UserRepository interface {
	CreateAccount(account *Account) error
	GetAccountByEmail(email string) (*Account, error)
}

// redis
type TokenRepository interface {
	SetAuthToken(login *Login) error
	GetAuthToken(login *Login) error
}
