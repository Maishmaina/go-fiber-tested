// main
package main

import (
    "github.com/gofiber/fiber/v2"
	"backend-dev-evaluation/controllers"
)

func main() {
    app := fiber.New()

	//basic route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	 // Register route
	 app.Post("/register", controllers.Register)

	 // Login route
	 app.Post("/login", controllers.Login)

    app.Listen(":3000")
}


/*
CREATE DATABASE IF NOT EXISTS chama_soft;
USE chama_soft;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
*/