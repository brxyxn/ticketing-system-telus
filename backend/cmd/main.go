package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./backend/build", fiber.Static{
		Index: "index.html",
	})
	app.Get("/api", handlerApi)

	app.Listen(":5000")
}

func handlerApi(c *fiber.Ctx) error {
	return c.SendString("Hello Golang!")
}
