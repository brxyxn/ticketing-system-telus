package users

type UserRepository interface {
	CreateAccount(account *Account) error
	GetAccountByEmail(email string) (*Account, error)
}
