package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

var users = []User{
	{ID: 1, Name: "James", Age: 10},
	{ID: 2, Name: "George", Age: 20},
	{ID: 3, Name: "John", Age: 30},
}

func getUser(c *fiber.Ctx) error {
	return c.JSON(users)
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + strconv.Itoa(users[0].ID) + " " + users[0].Name)
	})

	app.Get("/user", getUser)

	app.Listen(":8000")
}
