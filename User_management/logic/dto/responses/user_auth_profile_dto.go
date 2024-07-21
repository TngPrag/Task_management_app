package responses

import "github.com/go-playground/validator/v10"

type UserAuthProfileDto struct {
	UserID   string `json:"user_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (profile *UserAuthProfileDto) ValidateUserAuthProfileDto() error {
	validate := validator.New()
	if err := validate.Struct(profile); err != nil {
		return err
	}
	return nil
}
