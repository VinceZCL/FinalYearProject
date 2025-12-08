package param

import "github.com/go-playground/validator/v10"

type NewMember struct {
	UserID uint   `json:"userID" validate:"required"`
	TeamID uint   `json:"teamID" validate:"required"`
	Role   string `json:"role" validate:"required,max=50"`
}

func (r *NewMember) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
