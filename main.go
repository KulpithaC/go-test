package main

import (
	"github.com/KulpithaC/go-test/handlers"
	"github.com/KulpithaC/go-test/models"
	"github.com/KulpithaC/go-test/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	users := []models.User{}

	userService := services.NewUserService(&users)
	userHandler := handlers.NewUserHandler(userService)

	// api
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/users", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)

	app.Listen(":8080")
}
