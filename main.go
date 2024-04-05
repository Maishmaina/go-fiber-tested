// main
package main

import (
    "github.com/gofiber/swagger"
    "github.com/gofiber/fiber/v2"
	"backend-dev-evaluation/controllers"
    "backend-dev-evaluation/middleware"
)

//	@ChamaSoft		Fiber App API
//	@description	This is the API documentation for chamasoft API Fiber app
//	@version		1.0
//	@host			localhost:3000
//	@BasePath		/

func main() {
    app := fiber.New()

    //middleware to logs all request on this entry point
    app.Use(middleware.LoggerMiddleware)

    // Serve Swagger UI at /swagger/index.html
    app.Get("/swagger/*", swagger.HandlerDefault)
   

	//basic route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	 // Register route
	 app.Post("/register", controllers.Register)

	 // Login route
	 app.Post("/login", controllers.Login)

     //return all users 
     app.Get("/users",middleware.Authorization(middleware.RoleAdmin),controllers.AllUsers);

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