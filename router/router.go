package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yong509/go-fiber-playground/middlewares"
)

func Initalize(router *fiber.App) {
	router.Use(middlewares.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})
}
