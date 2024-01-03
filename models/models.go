package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"unique;email"`
	Password string `gorm:"password"`
	Limit    string  `gorm:"limit"`
}
type Grocery struct {
	gorm.Model
	Price int64  `gorm:"price"`
	Email string `gorm:"email"`
}
