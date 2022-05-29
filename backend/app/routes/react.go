package routes

import "github.com/gofiber/fiber/v2"

func ReactRoutes(app *fiber.App) {
	app.Static("/", "./build", fiber.Static{
		Index: "index.html",
	})
}
