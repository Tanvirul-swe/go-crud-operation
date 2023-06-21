package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 	  uint `gorm:"primaryKey"`
	Name	  string
	Email     string
	Password  string
	Image	 string
	Address   string
}
