package main

import "github.com/gofiber/fiber"

func main() {
	// Create new Fiber instance:
	app := fiber.New()

	// Create route on root path, "/":
	app.Get("/", func(c *fiber.Ctx) {
		c.JSON("Hello, World!")

	})

	// Start server on "localhost" with port "8080":
	app.Listen(8080)
}
