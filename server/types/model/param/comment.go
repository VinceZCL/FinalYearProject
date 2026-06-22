package param

import "github.com/go-playground/validator/v10"

type NewComment struct {
	UserID    uint   `json:"userID" validate:"required"`
	CheckinID uint   `json:"checkinID" validate:"required"`
	TeamID    uint   `json:"teamID" validate:"required"`
	Item      string `json:"item" validate:"required,max=255"`
}

func (r *NewComment) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
