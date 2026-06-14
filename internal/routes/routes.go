package routes

import (
	"github.com/Aditya-Sureka/Go_Backend_Task/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Running")
	})

	app.Post(
		"/users",
		userHandler.CreateUser,
	)

	app.Get(
		"/users/:id",
		userHandler.GetUser,
	)

	app.Get(
		"/users",
		userHandler.ListUsers,
	)
}
