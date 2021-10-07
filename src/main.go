package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	r := Response{
		Message: "Hellow World",
		Success: true,
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(r)
	})
	if err := app.Listen(":8081"); err != nil {
		fmt.Println(err)
		panic("Unable to run the app")
	}
}
