package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Create(app *fiber.App) {
	// assets
	app.Static("/static", "./static")

	// home
	app.Get("/", func(context *fiber.Ctx) error {
		return context.Render("index", fiber.Map{})
	})

}
