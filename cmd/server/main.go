package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Aditya-Sureka/Go_Backend_Task/config"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server Running")
	})

	app.Listen(":3000")
}