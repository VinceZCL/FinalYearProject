package models

import "gorm.io/gorm"

type CheckIn struct {
	gorm.Model
	UserID     uint   `gorm:"not null;index" json:"userID"`
	Type       string `gorm:"size:50;not null" json:"type"`
	Item       string `gorm:"size:255;not null" json:"item"`
	Jira       string `gorm:"size:100" json:"jira"`
	Visibility string `gorm:"size:10;default:'all';not null" json:"visibility"`
	TeamID     *uint  `gorm:"index" json:"teamID"`
	User       User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Team       Team   `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team"`
}
