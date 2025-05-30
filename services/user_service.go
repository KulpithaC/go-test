package services

import (
	"errors"

	"github.com/KulpithaC/go-test/models"
	"github.com/KulpithaC/go-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	// Check body
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("name, email, password are required")
	}

	// Check email duplicate
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return err
	}
	for _, u := range users {
		if u.Email == user.Email {
			return errors.New("email already exists")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Balance = 100

	return s.userRepo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
