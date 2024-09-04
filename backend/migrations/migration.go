package main

import (
	"gin-fleamarket/infra"
	"log"
	"gin-fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("failed to migrate")
	}
	log.Println("migration has been processed")
}
