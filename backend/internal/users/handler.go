package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserFiberHandler interface {
	RegisterAccount(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
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

// This function returns the authentication token to validate session
func (u *userFiberHandler) Authenticate(c *fiber.Ctx) error {
	value := c.Params("user_id")
	return c.JSON(value)
}
