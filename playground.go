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

func getAllUser(c *fiber.Ctx) error {
	return c.JSON(users)
}

func findUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			return c.JSON(user)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "create user unsuccessfully",
			"status":  fiber.StatusBadRequest,
		})
	}
	user.ID = len(users) + 1
	users = append(users, *user)

	return c.JSON(users)
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + strconv.Itoa(users[0].ID) + " " + users[0].Name)
	})

	app.Get("/user", getAllUser)
	app.Get("/user/:id", findUser)
	app.Post("/create/user", createUser)

	app.Listen(":8000")
}
