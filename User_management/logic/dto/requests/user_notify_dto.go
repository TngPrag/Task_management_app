package requests

import "github.com/go-playground/validator/v10"

type UserNotifyDto struct {
	User_id     string `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Deadline    string `json:"deadline" validate:"required"`
}

func (notify *UserNotifyDto) ValidateUserNotifyDto() error {
	validate := validator.New()

	if err := validate.Struct(notify); err != nil {
		return err
	}
	return nil
}
