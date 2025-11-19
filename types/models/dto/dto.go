package dto

type User struct {
	UserID uint   `json:"userID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Type   string `json:"type"`
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
	TeamID      uint   `json:"teamID"`
	TeamName    string `json:"team_name"`
	CreatorID   uint   `json:"creatorID"`
	CreatorName string `json:"creator_name"`
}
