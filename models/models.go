package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" gorm:"text;not null;default:null`
	Price string `json:"price" gorm:"text;not null;default:null`
}
