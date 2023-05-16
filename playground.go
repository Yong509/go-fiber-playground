package main

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	users := []User{
		{ID: 1, Name: "Test", Age: 10},
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + strconv.Itoa(users[0].ID) + " " + users[0].Name)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		returnUsers, _ := json.Marshal(users)
		return c.SendString(string(returnUsers))
	})

	app.Listen(":8000")
}
