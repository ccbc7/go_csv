package models

import "gorm.io/gorm"

type Csv struct {
	gorm.Model
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	PhoneNumber string
	Address     string
	City        string
	State       string
	ZipCode     string
	Country     string
}
