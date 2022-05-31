package routes

import "github.com/gofiber/fiber/v2"

func ReactRoutes(app *fiber.App) {
	app.Static("/", "./build", fiber.Static{
		Index: "index.html",
	})

	app.Static("/login", "./build", fiber.Static{
		Index: "index.html",
	})

	app.Static("/register", "./build", fiber.Static{
		Index: "index.html",
	})

	app.Static("/tickets", "./build", fiber.Static{
		Index: "index.html",
	})
}
