package services

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"tyres.kz/internal/models"
	"tyres.kz/internal/repositories"
)

type UserServiceInterface interface {
	Create(user *models.User) (int, error)
	GetByID(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByPhone(phone string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
	logger         *log.Logger
}

func NewUserService(userRepository repositories.UserRepositoryInterface, logger *log.Logger) (UserServiceInterface, error) {
	if userRepository == nil {
		return nil, fmt.Errorf("userRepository is nil")
	}

	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &UserService{userRepository: userRepository, logger: logger}, nil
}

func (service *UserService) Create(user *models.User) (int, error) {
	passwordByte := []byte(user.HashedPassword)

	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)

	if err != nil {
		service.logger.Printf("Error hashing password: %v", err)
		return 0, err
	}

	user.HashedPassword = string(hash)

	return service.userRepository.Create(user)
}

func (service *UserService) GetByID(id int) (*models.User, error) {
	return service.userRepository.GetByID(id)
}

func (service *UserService) GetByUsername(username string) (*models.User, error) {
	return service.userRepository.GetByUsername(username)
}

func (service *UserService) GetByEmail(email string) (*models.User, error) {
	return service.userRepository.GetByEmail(email)
}

func (service *UserService) GetByPhone(phone string) (*models.User, error) {
	return service.userRepository.GetByPhone(phone)
}

func (service *UserService) Update(user *models.User) error {
	return service.userRepository.Update(user)
}

func (service *UserService) Delete(id int) error {
	return service.userRepository.Delete(id)
}
