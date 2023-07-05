

package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryId uint `gorm:"primaryKey"`
	Name string `json:"name"`
}