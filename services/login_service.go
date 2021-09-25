package services

import (
	"fmt"

	"github.com/a180024/golang-api/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(username string, password string) (bool, error)
}

type loginService struct {
	userRepository models.UserRepository
}

func NewLoginService(userRepository models.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (loginService *loginService) Login(username string, password string) (bool, error) {
	// Fetch user details from DB
	user, err := loginService.userRepository.FindOneByUserName(username)
	if err != nil {
		return false, err
	}
	fmt.Println("password", user.Password)

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
