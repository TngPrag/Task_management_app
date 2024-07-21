package requests

import "github.com/go-playground/validator/v10"

type LoginDto struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string `json:"password" valdiate:"required,min=8"`
}

func (login *LoginDto) ValidateLoginRequestDto() error {
	validate := validator.New()
	if err := validate.Struct(login); err != nil {
		return err
	}
	return nil
}
