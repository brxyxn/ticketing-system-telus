package users

import (
	"errors"
	"os"

	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(email, password string) (*Login, error)
	Register(account *Account) error

	SetAuthToken(login *Login) error
}

type userService struct {
	dbRepo    UserRepository
	cacheRepo TokenRepository
	key       []byte
}

func NewUserService(databaseRepository UserRepository, cacheRepository TokenRepository) UserService {
	key := os.Getenv("SECRET_KEY")
	return &userService{databaseRepository, cacheRepository, []byte(key)}
}

func (s *userService) Register(account *Account) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(account.User.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Error(err)
		return err
	}

	account.User.Password = string(passwordHash)

	err = s.dbRepo.CreateAccount(account)
	if err != nil {
		u.Log.Error(err)
		return err
	}

	account.User.Password = ""

	return nil
}

func (s *userService) Login(email, password string) (*Login, error) {
	account, err := s.dbRepo.GetAccountByEmail(email)
	if err != nil {
		u.Log.Error(err)
		return &Login{LoggedIn: false}, err
	}
	if account == nil {
		return &Login{LoggedIn: false}, errors.New("Invalid email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.User.Password), []byte(password)); err != nil {
		return &Login{LoggedIn: false}, errors.New("Invalid password")
	}

	return &Login{
		LoggedIn: true,
	}, nil
}

func (s *userService) SetAuthToken(login *Login) error {
	return s.cacheRepo.SetAuthToken(login)
}
