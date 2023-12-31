package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}
type Grocery struct {
	gorm.Model
	Price int64  `gorm:"price"`
	Email string `gorm:"email"`
}
