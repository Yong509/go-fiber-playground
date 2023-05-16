package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("receive parameter " + c.Params("value"))
	})

	app.Get("/more/*", func(c *fiber.Ctx) error {
		return c.SendString("more parameters " + c.Params("*"))
	})

	app.Listen(":8000")
}
