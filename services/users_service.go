package services

import (
	"github.com/a180024/golang-api/dto"
	"github.com/a180024/golang-api/models"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository models.UserRepository
}

type UserService interface {
	CreateUser(userDto dto.UserDto) error
}

func NewUserService(userRepository models.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}

}

func (userService *userService) CreateUser(userDto dto.UserDto) error {
	if err := userDto.Validate(); err != nil {
		return err
	}

	// Password encryption
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), 14)
	if err != nil {
		return err
	}
	userDto.Password = string(pwSlice[:])

	// Save in DB
	err = userService.userRepository.Save(userDto)
	if err != nil {
		return err
	}

	return nil
}
