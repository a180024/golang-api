package services

import (
	"github.com/a180024/nft_api/dto/users"
	"github.com/a180024/nft_api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.ErrResponse) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// password encryption
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("Failed to encrypt the password")
	}

	user.Password = string(pwSlice[:])

	return &user, nil
}
