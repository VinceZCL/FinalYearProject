package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Email    string `gorm:"size:100;not null" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	Type     string `gorm:"size:50;default:'user'" json:"type"`
	Status   string `gorm:"size:50;default:'active'" json:"status"`
}
