package middleware

import (
    "github.com/gofiber/fiber/v2"
)

// Define roles
const (
    RoleAdmin = "admin"
    RoleUser  = "user"
)

// Authorization
func Authorization(role string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Check user role
        userRole := "user" 

        // Check if the user is authorized to access the route based on role
        if userRole == role || userRole == RoleAdmin {
            return c.Next()
        }

        // User is not authorized
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized access"})
    }
}