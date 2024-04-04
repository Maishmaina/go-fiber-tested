package controllers

import (
    "github.com/gofiber/fiber/v2"
    "backend-dev-evaluation/models"
   
)

func Register(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing request"})
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error hashing password"})
    }

    user.Password = hashedPassword

    // Save the user to the database
   
    // db.Exec("INSERT INTO users (name, email,password) VALUES (?, ?)", user.Name, user.Email,user.Password)

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing request"})
    }

    // Retrieve user from the database

    // db.Get(&user, "SELECT * FROM users WHERE email = ?", user.Email)

    // Check if the user exists and verify the password
    if user.ID == 0 || !utils.CheckPasswordHash(user.Password, user.Password) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error generating JWT token"})
    }

    return c.JSON(fiber.Map{"token": token})
}