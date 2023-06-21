package model

import (
	"gorm.io/gorm"
)

type UserFevoriteBook struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	UserId uint `gorm:"foreignKey:ID"`
	BookId uint `gorm:"foreignKey:ID"`
	
}