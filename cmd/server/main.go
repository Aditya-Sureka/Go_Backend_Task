package main

import (
"log"


"github.com/Aditya-Sureka/Go_Backend_Task/config"
"github.com/gofiber/fiber/v2"


)

func main() {


db := config.ConnectDB()
defer db.Close()

app := fiber.New()

app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Server Running")
})

log.Fatal(app.Listen(":3000"))


}
