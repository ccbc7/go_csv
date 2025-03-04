package config

import (
	"log"

	"github.com/joho/godotenv"
)

// envファイルの読み込み
func Initialize() {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
