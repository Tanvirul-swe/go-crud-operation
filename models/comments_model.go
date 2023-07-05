package model

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	BookId uint `gorm:"foreignKey:ID"`
	UserId uint `gorm:"foreignKey:ID"`
	Comment string
}