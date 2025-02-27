package models

import (
	"time"
)

// Item represents the item model
type Item struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name" gorm:"not null"`
	Price       uint       `json:"price" gorm:"not null"`
	Description string     `json:"description"`
	SoldOut     bool       `json:"sold_out" gorm:"not null,default:false"`
	UserID      uint       `json:"user_id" gorm:"not null"`
}
