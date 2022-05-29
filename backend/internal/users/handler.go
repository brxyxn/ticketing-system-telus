package users

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserFiberHandler interface {
	RegisterAccount(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

type userFiberHandler struct {
	service UserService
}

func NewUserFiberHandler(service UserService) UserFiberHandler {
	return &userFiberHandler{service}
}

// This function registers a new user, and returns the account details.
// If the account is for a customer, company name will be created.
// If the account is for an agent, a team will be assigned.
// By default, the account is created with the role of customer.
// And the profile is created at the same time.
func (u *userFiberHandler) RegisterAccount(c *fiber.Ctx) error {
	var account Account
	err := c.BodyParser(&account)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = u.service.Register(&account)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(account)
}

type Locals struct {
	User *jwt.Token
	Ctx  *fiber.Ctx
}

// This function returns the authentication token to validate session
func (u *userFiberHandler) Authenticate(c *fiber.Ctx) error {
	var login Login
	var err error

	err = c.BodyParser(&login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if login.Email == "" || login.Password == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	auth, err := u.service.Login(login.Email, login.Password)
	if err != nil || !auth.LoggedIn {
		return c.Status(fiber.StatusUnauthorized).JSON(errors.New("Invalid email or password"))
	}

	var local Locals
	local.Ctx = c
	auth.Token, err = local.authenticate(auth)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = u.service.SetAuthToken(auth)

	log.Println("token", auth)

	return c.JSON(fiber.Map{"token": auth.Token})
}

func (c *Locals) authenticate(login *Login) (string, error) {
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

func (u *userFiberHandler) GetUser(c *fiber.Ctx) error {
	var login Login
	err := u.service.GetAuthToken(&login)
	if err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(errors.New("Invalid token"))
	}
	tokenString := c.GetReqHeaders()["Authorization"]
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	log.Println("tkn:", login.Token, "token:", tokenString)

	if login.Token != tokenString {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.SendString("Hello")
}
