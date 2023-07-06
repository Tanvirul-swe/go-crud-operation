package model

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey"`
	Name        string   `json:"name"`
	Author_Name string   `json:"author_name"`
	Category    int      `json:"category_id"`
	Discription string   `json:"discription"`
	Image       string   `json:"image"`
	Book_Url    string   `json:"book_url"`
	User_Id     int      `gorm:"foreignKey:User_Id"`
	Status      int      `json:"status"`
	Categorys   Category `gorm:"foreignKey:Category"`
}
