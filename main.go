package main

import (
	"log"

	db "github.com/KulpithaC/go-test/database"
	"github.com/KulpithaC/go-test/handlers"
	"github.com/KulpithaC/go-test/models"
	"github.com/KulpithaC/go-test/repository"
	"github.com/KulpithaC/go-test/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	var users []models.User

	if err := db.ConnectDB(); err != nil {
		log.Fatal("Can not start application due to database connection issue")
	}

	if err := db.DB.AutoMigrate(&users); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	userRepo := repository.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// api
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/users", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)

	app.Listen(":8080")
}
