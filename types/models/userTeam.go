package models

type UserTeam struct {
	UserID uint   `gorm:"primaryKey;autoIncrement:false" json:"userID"`
	TeamID uint   `gorm:"primaryKey;autoIncrement:false" json:"teamID"`
	Role   string `gorm:"size:50;not null;default:'member'" json:"role"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Team   Team   `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team"`
}
