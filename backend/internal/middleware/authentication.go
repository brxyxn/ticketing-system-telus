package middleware

import (
	"strings"

	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type TokenHandler interface {
	Validate(c *fiber.Ctx) error
}

type tokenHandler struct {
	service UserService
	store   *session.Store
}

func NewTokenHandler(service UserService, store *session.Store) TokenHandler {
	return &tokenHandler{service, store}
}

func (a *tokenHandler) Validate(c *fiber.Ctx) error {
	u.Log.Info("Middleware: Authenticating user")
	var login Login
	c.BodyParser(&login)

	// getting token from headers
	tokenString := c.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"response": "missing token"})
	}

	login.IP = c.IP()

	// validating token
	a.service.GetAuthToken(&login)
	if login.Token != tokenString {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"response": "invalid token"})
	}

	c.Next()
	return nil
}
