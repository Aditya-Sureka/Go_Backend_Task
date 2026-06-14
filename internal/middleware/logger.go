package middleware

import (
"time"

"github.com/Aditya-Sureka/Go_Backend_Task/internal/logger"
"github.com/gofiber/fiber/v2"
"go.uber.org/zap"

)

func RequestLogger() fiber.Handler {

return func(c *fiber.Ctx) error {

	start := time.Now()

	err := c.Next()

	logger.Log.Info(
		"http request",
		zap.String("request_id", c.Locals("requestID").(string)),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Duration("duration", time.Since(start)),
		zap.Int("status", c.Response().StatusCode()),
	)

	return err
}

}