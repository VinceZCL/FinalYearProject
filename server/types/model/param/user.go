package param

import "github.com/go-playground/validator/v10"

type NewUser struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Type     string `json:"type" validate:"omitempty,oneof=user admin,max=50"`
}

func (r *NewUser) Validate() error {
	v := validator.New()
	return v.Struct(r)
}

type UpdateUser struct {
	UserID          uint    `json:"userID" validate:"required"`
	Name            string  `json:"name" validate:"required,max=100"`
	Email           string  `json:"email" validate:"required,email,max=100"`
	CurrentPassword string  `json:"current_password" validate:"required,max=100"`
	NewPassword     *string `json:"new_password" validate:"max=100"`
}

func (r *UpdateUser) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
