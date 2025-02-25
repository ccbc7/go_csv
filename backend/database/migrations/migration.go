package main

import (
	"log"
	"project/infra"
	"project/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.Csv{}); err != nil {
		panic("failed to migrate")
	}
	log.Println("migration has been processed")
}
