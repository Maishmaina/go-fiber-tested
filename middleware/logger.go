package middleware

import (
    "github.com/gofiber/fiber/v2"
    "log"
    "time"
)

func LoggerMiddleware(c *fiber.Ctx) error {
    // Log
    log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), c.Method(), c.Path())

   
    return c.Next()
}