package dto

import "github.com/go-playground/validator/v10"

type UserDto struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
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

func (userDto *UserDto) Validate() error {
	v := validator.New()
	if err := v.Struct(userDto); err != nil {
		return err
	}
	return nil
}
