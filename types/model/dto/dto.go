package dto

import "time"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Member struct {
	UserID   uint   `json:"userID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	TeamID   uint   `json:"teamID"`
	TeamName string `json:"team_name"`
	Role     string `json:"role"`
}

type Team struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	CreatorID   uint   `json:"creatorID"`
	CreatorName string `json:"creator_name"`
}

type CheckIn struct {
	CheckInID  uint      `json:"id"`
	Type       string    `json:"type"`
	Item       string    `json:"item"`
	Jira       string    `json:"jira"`
	Visibility string    `json:"visibility"`
	TeamID     *uint     `json:"teamID"`
	UserID     uint      `json:"userID"`
	Username   string    `json:"user_name"`
	CreatedAt  time.Time `json:"created_at"`
}
