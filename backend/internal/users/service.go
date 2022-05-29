package users

import (
	"errors"
	"os"
	"time"

	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(username, password string) (*Login, error)
	Register(account *Account) error
}

type userService struct {
	repo UserRepository
	key  []byte
}

func NewUserService(repository UserRepository) UserService {
	key := os.Getenv("SECRET_KEY")
	return &userService{repository, []byte(key)}
}

func (s *userService) Register(account *Account) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(account.User.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Error(err)
		return err
	}

	account.User.Password = string(passwordHash)

	err = s.repo.CreateAccount(account)
	if err != nil {
		u.Log.Error(err)
		return err
	}

	account.User.Password = ""

	return nil
}

func (s *userService) Login(email, password string) (*Login, error) {
	account, err := s.repo.GetAccountByEmail(email)
	if err != nil {
		u.Log.Error(err)
		return nil, err
	}
	if account == nil {
		return nil, errors.New("Invalid email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.User.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid password")
	}

	token, err := s.getToken(account)
	if err != nil {
		return nil, err
	}

	return &Login{
		Email: email,
		Token: token,
	}, nil
}

func (s *userService) getToken(account *Account) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = account.User.UserID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString(s.key)
}
