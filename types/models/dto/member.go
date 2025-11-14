package dto

type Member struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	TeamName string `json:"team_name"`
	Role     string `json:"role"`
}
