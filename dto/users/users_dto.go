package users

import (
	"strings"

	"github.com/a180024/nft_api/utils/errors"
)

type UserDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *UserDto) Validate() *errors.ErrResponse {
	user.UserName = strings.TrimSpace(user.UserName)
	user.Password = strings.TrimSpace(user.Password)
	if user.UserName == "" {
		return errors.NewBadRequestError("Invalid Username")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
