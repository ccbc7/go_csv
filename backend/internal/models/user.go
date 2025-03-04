package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	// ユーザーが削除された場合、そのユーザーのアイテムも削除される
	Items []Item `gorm:"constraint:OnDelete:CASCADE"`
}
