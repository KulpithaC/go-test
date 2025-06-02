package handlers

import (
	"strconv"

	"github.com/KulpithaC/go-test/models"
	"github.com/KulpithaC/go-test/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	response := []UserResponse{}
	for _, user := range users {
		response = append(response, UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Balance: user.Balance,
		})
	}

	return c.JSON(response)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	response := UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Balance: user.Balance,
	}

	return c.JSON(response)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.userService.CreateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Balance: user.Balance,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	_, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
	})
}
