package dto

import (
	"errors"
	"strings"
)

type UserDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *UserDto) Validate() error {
	user.UserName = strings.TrimSpace(user.UserName)
	user.Password = strings.TrimSpace(user.Password)
	if user.UserName == "" {
		return errors.New("Invalid Username")
	}
	if user.Password == "" {
		return errors.New("Invalid password")
	}
	return nil
}
