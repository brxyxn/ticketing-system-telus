package users

import (
	"errors"
	"net/http"
	"time"

	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler interface {
	RegisterAccount(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

type userHandler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return &userHandler{service}
}

// This function registers a new user, and returns the account details.
// If the account is for a customer, company name will be created.
// If the account is for an agent, a team will be assigned.
// By default, the account is created with the role of customer.
// And the profile is created at the same time.
func (a *userHandler) RegisterAccount(c *fiber.Ctx) error {
	u.Log.Info("Registering new user")
	var account Account
	err := c.BodyParser(&account)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = a.service.Register(&account)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(account)
}

// This function returns the authentication token to validate session
func (a *userHandler) Login(c *fiber.Ctx) error {
	u.Log.Info("Authenticating user")
	var login Login
	var err error

	err = c.BodyParser(&login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if login.Email == "" || login.Password == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	auth, err := a.service.Login(login.Email, login.Password)
	if err != nil || !auth.LoggedIn {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.New("Invalid email or password"))
	}

	auth.Token, err = authenticate(auth)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	auth.Email = login.Email
	auth.IP = c.IP()
	err = a.service.SetAuthToken(auth)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(errors.New("error setting auth token"))
	}

	return c.JSON(fiber.Map{"token": auth.Token})
}

func authenticate(login *Login) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"email": login.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *userHandler) GetUser(c *fiber.Ctx) error {
	u.Log.Info("Getting user")

	return c.SendString("Hello")
}
