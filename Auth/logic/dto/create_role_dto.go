package dto

import "github.com/go-playground/validator/v10"

type CreateRoleDto struct {
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role"  validate:"required"`
}

func (role *CreateRoleDto) ValidateCreateRoleDto() error {
	validate := validator.New()
	if err := validate.Struct(role); err != nil {
		return err
	}
	return nil
}
