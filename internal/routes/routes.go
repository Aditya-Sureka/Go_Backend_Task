package routes

import (
	"github.com/Aditya-Sureka/Go_Backend_Task/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {

	app.Post(
		"/users",
		userHandler.CreateUser,
	)
}