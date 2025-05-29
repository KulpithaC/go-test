package services

import (
	"errors"

	"github.com/KulpithaC/go-test/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users *[]models.User
}

func NewUserService(users *[]models.User) *UserService {
	return &UserService{
		users: users,
	}
}

func (s *UserService) GetAllUsers() []models.User {
	return *s.users
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	for _, user := range *s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserService) CreateUser(user *models.User) error {
	// Check body
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("name, email, password are required")
	}

	// Check email
	for _, u := range *s.users {
		if u.Email == user.Email {
			return errors.New("email already exists")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set user data
	user.Password = string(hashedPassword)
	user.ID = len(*s.users) + 1
	user.Balance = 100

	*s.users = append(*s.users, *user)

	return nil
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	// Check body
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	for _, user := range *s.users {
		if user.Email == email {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return nil, errors.New("invalid email or password")
			}
			return &user, nil
		}
	}
	return nil, errors.New("invalid email or password")
}
