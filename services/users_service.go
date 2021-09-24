package services

import (
	"fmt"

	"github.com/a180024/nft_api/dto/users"
	"github.com/a180024/nft_api/models"
	"github.com/a180024/nft_api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository models.UserRepository
}

type UserService interface {
	CreateUser(userDto users.UserDto) (*users.UserDto, *errors.ErrResponse)
}

func newUserService(userRepository models.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}

}

func (userService *userService) CreateUser(userDto users.UserDto) (*users.UserDto, *errors.ErrResponse) {
	if err := userDto.Validate(); err != nil {
		return nil, err
	}

	// password encryption
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("Failed to encrypt the password")
	}
	userDto.Password = string(pwSlice[:])

	// save in DB
	user, err := userService.userRepository.Save(userDto)

	fmt.Println(user)

	return &userDto, nil
}
