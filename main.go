package main

import (
	"fiber-template/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	isProd := getEnvironmentMode()

	viewEngine := html.New("./views", ".tmpl")

	// reload the templates on each render when in dev mode
	viewEngine.Reload(!isProd)

	app := fiber.New(fiber.Config{
		Views:       viewEngine,
		ViewsLayout: "layouts/default",
	})

	routes.Create(app)

	// 404 Config
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{})
	})

	log.Fatal(app.Listen(":8080"))
}

/**
Privates
=========
*/

func getEnvironmentMode() bool {
	_ = godotenv.Load() // loads the .env if it exists

	isProd, err := strconv.ParseBool(os.Getenv("IS_PROD"))

	if err != nil {
		log.Println("Failed to parse 'IS_PROD' environment variable")
	}

	return isProd
}
