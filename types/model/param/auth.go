package param

import "github.com/go-playground/validator/v10"

type Login struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

func (r *Login) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
