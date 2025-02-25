package seeders

import (
	"log"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	// 一度データを削除
	if err := db.Exec("DELETE FROM items").Error; err != nil {
		log.Fatalf("アイテムの削除に失敗しました: %v", err)
	}
	log.Println("アイテムの削除が完了しました")

	seeds := []func(*gorm.DB) error{
		SeedItems,
	}

	for _, seed := range seeds {
		if err := seed(db); err != nil {
			log.Fatalf("failed to seed: %v", err)
		}
	}

	log.Println("All seeding completed successfully")
}
