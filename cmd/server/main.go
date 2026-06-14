package main

import (
"log"


"github.com/Aditya-Sureka/Go_Backend_Task/config"
"github.com/Aditya-Sureka/Go_Backend_Task/db/sqlc"
"github.com/Aditya-Sureka/Go_Backend_Task/internal/handler"
"github.com/Aditya-Sureka/Go_Backend_Task/internal/repository"
"github.com/Aditya-Sureka/Go_Backend_Task/internal/routes"
"github.com/Aditya-Sureka/Go_Backend_Task/internal/service"
"github.com/gofiber/fiber/v2"
"github.com/Aditya-Sureka/Go_Backend_Task/internal/logger"


)

func main() {

logger.Init()
defer logger.Log.Sync()

db := config.ConnectDB()
defer db.Close()

queries := sqlc.New(db)

userRepo := repository.NewUserRepository(queries)

userService := service.NewUserService(userRepo)

userHandler := handler.NewUserHandler(userService)

app := fiber.New()

routes.Setup(app, userHandler)

log.Fatal(app.Listen(":3000"))


}
