package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null" json:"name"`
	Email    string `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	Type     string `gorm:"size:50;default:'user'" json:"type"`
	Status   string `gorm:"size:50;default:'active'" json:"status"`
}

type Team struct {
	gorm.Model
	Name      string `gorm:"size:100;not null;index:idx_name_creator,unique" json:"name"`
	CreatorID uint   `gorm:"not null;index:idx_name_creator,unique" json:"creatorID"`
	Creator   User   `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"creator"`
}

type UserTeam struct {
	UserID uint   `gorm:"primaryKey;autoIncrement:false" json:"userID"`
	TeamID uint   `gorm:"primaryKey;autoIncrement:false" json:"teamID"`
	Role   string `gorm:"size:50;not null;default:'member'" json:"role"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Team   Team   `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team"`
}

type CheckIn struct {
	gorm.Model
	UserID     uint    `gorm:"not null;index" json:"userID"`
	Type       string  `gorm:"size:50;not null" json:"type"`
	Item       string  `gorm:"size:255;not null" json:"item"`
	Jira       *string `gorm:"size:100" json:"jira"`
	Visibility string  `gorm:"size:10;default:'all';not null" json:"visibility"`
	TeamID     *uint   `gorm:"index" json:"teamID"`
	User       User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Team       *Team   `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team"`
}
