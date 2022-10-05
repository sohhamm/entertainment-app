package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sohhamm/entertainment-app/database"
)

func main() {
	// connect to the database

	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Entertainment App API")
	})

	app.Listen(":9000")
}
