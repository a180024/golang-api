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

type UserIdDto struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UserResponseDto struct {
	ID        string `json:"user_id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
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
