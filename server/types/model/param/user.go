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
