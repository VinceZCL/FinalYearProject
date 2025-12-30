package param

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type NewCheckIn struct {
	UserID     uint    `json:"userID" validate:"required"`
	Type       string  `json:"type" validate:"required,oneof=yesterday today blockers,max=50"`
	Item       string  `json:"item" validate:"required,max=255"`
	Jira       *string `json:"jira" validate:"omitempty,max=100"`
	Visibility string  `json:"visibility" validate:"required,oneof=all private team,max=10"`
	TeamID     *uint   `json:"teamID" validate:"omitempty"`
}

type BulkCheckIn struct {
	CheckIns []NewCheckIn `json:"checkIns" validate:"required,dive"`
}

func (r *NewCheckIn) Validate() error {
	v := validator.New()
	return v.Struct(r)
}

// Validate method for BulkCheckIn to enforce rules for the batch
func (b *BulkCheckIn) Validate() error {
	v := validator.New()

	// First validate the fields inside each NewCheckIn
	if err := v.Struct(b); err != nil {
		return err
	}

	// Custom validation to ensure there's at least one "yesterday" and one "today"
	var hasYesterday, hasToday bool

	// Loop through each check-in and check its type
	for _, checkIn := range b.CheckIns {
		if checkIn.Type == "yesterday" {
			hasYesterday = true
		}
		if checkIn.Type == "today" {
			hasToday = true
		}
	}

	// If there's no "yesterday" check-in, return an error
	if !hasYesterday {
		return errors.New("at least one 'yesterday' check-in is required")
	}

	// If there's no "today" check-in, return an error
	if !hasToday {
		return errors.New("at least one 'today' check-in is required")
	}

	return nil
}
