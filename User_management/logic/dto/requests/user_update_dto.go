package requests

import "github.com/go-playground/validator/v10"

type UserCredentialUpdateDto struct {
	UserName string `json:"user_name"  validate:"required"`
	Password string `json:"password"   validate:"required"`
}

func (change *UserCredentialUpdateDto)UpdateUserCredentialDto() error{
	validate := validator.New()
	err := validate.Struct(change)
	if err != nil{
		return err
	}
	return nil
}