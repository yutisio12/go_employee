package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	BadgeId  uint `gorm:"primaryKey" json:"badge_id"`
	Name     string `json:"name"`
	Badge    int    `json:"badge"`
	Mail     string `json:"mail"`
}
