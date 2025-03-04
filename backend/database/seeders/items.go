package seeders

import (
	"project/internal/models"

	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB) error {
	items := []models.Item{
		{Name: "Item1", Price: 100, Description: "Description1", SoldOut: false, UserID: 1},
		{Name: "Item2", Price: 200, Description: "Description2", SoldOut: false, UserID: 2},
		{Name: "Item3", Price: 300, Description: "Description3", SoldOut: false, UserID: 1},
	}

	for _, item := range items {
		if err := db.Create(&item).Error; err != nil {
			return err
		}
	}
	return nil
}
