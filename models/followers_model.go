package model

import "gorm.io/gorm"

type Followers struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	UserId uint `gorm:"foreignKey:ID"`
	FollowerId uint `gorm:"foreignKey:ID"`
}