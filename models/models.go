package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"unique;email"`
	Password string `gorm:"password"`
	Limit    string `gorm:"limit"`
}
type Grocery struct {
	gorm.Model
	Expense string    `gorm:"expense"`
	Email   string    `gorm:"email"`
	Date    time.Time `gorm:"type:date"`
}

type Medicine struct {
	gorm.Model
	Expense string    `gorm:"expense"`
	Email   string    `gorm:"email"`
	Date    time.Time `gorm:"type:date"`
}

type Transportation struct {
	gorm.Model
	Expense string    `gorm:"expense"`
	Email   string    `gorm:"email"`
	Date    time.Time `gorm:"type:date"`
}

type HomeMaintanance struct {
	gorm.Model
	Expense string    `gorm:"expense"`
	Email   string    `gorm:"email"`
	Date    time.Time `gorm:"type:date"`
}

type PDFDetails struct {
	ID        int        `json:"ID"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`
	Expense   string     `json:"Expense"`
	Email     string     `json:"Email"`
	Date      time.Time  `json:"Date"`
}
