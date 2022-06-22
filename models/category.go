package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	ParentID *uint     `json:"parent_id" gorm:"default:null"`
	Category *Category `gorm:"foreignKey:ParentID"`
}
