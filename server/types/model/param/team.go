package param

import "github.com/go-playground/validator/v10"

type NewTeam struct {
	Name      string `json:"name" validate:"required,max=100"`
	CreatorID uint   `json:"creatorID" validate:"required"`
}

func (r *NewTeam) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
