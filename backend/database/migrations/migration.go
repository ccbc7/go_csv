package main

import (
	"log"
	"project/config"
	"project/database"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	config.Initialize()
	db := database.SetupDB()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202503011700",
			Migrate: func(tx *gorm.DB) error {
				// 新しいカラムの追加
				type Item struct {
					NewColumn string `gorm:"type:varchar(100)"`
				}
				return tx.AutoMigrate(&Item{})
			},
			Rollback: func(tx *gorm.DB) error {
				// カラムの削除
				return tx.Migrator().DropColumn("items", "new_column")
			},
		},
		{
			ID: "202503011701",
			Migrate: func(tx *gorm.DB) error {
				// Userテーブルの作成
				type User struct {
					ID       uint   `gorm:"primaryKey"`
					Name     string `gorm:"type:varchar(100)"`
					Email    string `gorm:"type:varchar(100);unique"`
					Password string `gorm:"type:varchar(255)"`
				}
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Migration did run successfully")
}
