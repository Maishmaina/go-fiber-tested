package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsersHandler(t *testing.T) {
	app := setupTestApp()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Add more assertions here as needed
}

func setupTestApp() *fiber.App {
	app := fiber.New()

	controllers := &UserControllerMock{} // Use a mock UserController for testing
	app.Get("/users", controllers.AllUsers)

	return app
}

// Define a mock UserController for testing purposes
type UserControllerMock struct{}

func (uc *UserControllerMock) AllUsers(c *fiber.Ctx) error {
	// Mock implementation for GetAllUsers handler
	return c.SendString("Mocked response")
}