package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name      string `gorm:"size:100;not null;index" json:"name"`
	CreatorID uint   `gorm:"not null;index" json:"creatorID"`
	Creator   User   `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"creator"`
}
