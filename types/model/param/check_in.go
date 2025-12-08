package param

import "github.com/go-playground/validator/v10"

type NewCheckIn struct {
	UserID     uint    `json:"userID" validate:"required"`
	Type       string  `json:"type" validate:"required,oneof=yesterday today blockers,max=50"`
	Item       string  `json:"item" validate:"required,max=255"`
	Jira       *string `json:"jira" validate:"omitempty,max=100"`
	Visibility string  `json:"visibility" validate:"required,oneof=all private team,max=10"`
	TeamID     *uint   `json:"teamID" validate:"omitempty"`
}

func (r *NewCheckIn) Validate() error {
	v := validator.New()
	return v.Struct(r)
}
