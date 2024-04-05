package controllers

import (
    "github.com/gofiber/fiber/v2"
    "backend-dev-evaluation/models"
    "backend-dev-evaluation/utils"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
   
)

var db *sql.DB

func init() {
    // db connection
    db, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/chama_soft")
    if err := db.Ping(); err != nil {
        panic(err)
    }
}

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
    _, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving user"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing request"})
    }

    // Retrieve user from the database

    row := db.QueryRow("SELECT id, password FROM users WHERE email = ?", user.Email)

    var storedPassword string
    var userID int
    if err := row.Scan(&userID, &storedPassword); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Emails or Password"})
    }

    // Check if the user exists and verify the password
    if !utils.CheckPasswordHash(user.Password, user.Password) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Email or Passwords"})
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error generating JWT token"})
    }

    return c.JSON(fiber.Map{"token": token})
}

//fetch all users from the db
func AllUsers(c *fiber.Ctx)error{

    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching users"})
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error scanning users"})
        }
        users = append(users, user)
    }

    // Return users
    return c.JSON(users)
}