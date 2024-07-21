package dto

import "github.com/go-playground/validator/v10"


type CreatePolicyDto struct {
	Subject   string `json:"sub"`
	Object    string  `json:"object"`
	Action    string   `json:"action"`
}

func (policy *CreatePolicyDto) ValidateCreatePolicyDto()error{
	validate := validator.New()
	if err := validate.Struct(policy); err!= nil{
		return err
	}
	return nil
}
