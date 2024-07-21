package requests

import "github.com/go-playground/validator/v10"

type CreateUserDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	UserName  string `json:"user_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	//Role string `json:"role" validate:"required"`
}

func (user *CreateUserDto) ValidateCreateUserDto() error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil

}
