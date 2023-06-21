package model

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string 
	Author_Name string 
	Category int `gorm:"foreignKey"`
	Discription string
	Image string
	Book_Url string
	User_Id int `gorm:"foreignKey:ID"`
	Status int

	
}