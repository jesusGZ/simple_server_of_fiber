package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jesusGZ/simple_server_of_fiber/pkg/route"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "API",
		})
	})

	route.Expenses(app)
	app.Listen(":3000")
}
