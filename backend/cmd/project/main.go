package main

import (
	"flag"
	"fmt"
	"project/config"
	"project/database"
	"project/database/seeders"
	"project/internal/router"
)

// @title			Swagger Example API
// @version		1.0
// @description	This is a sample server for a pet store.
// @BasePath		/api/v1
func main() {
	seed := flag.Bool("seed", false, "Run the database seeders")
	flag.Parse()

	// 初期化(環境変数の読み込み)
	config.Initialize()

	// DB接続
	db := database.SetupDB()

	if *seed {
		seeders.SeedAll(db)
		fmt.Println("Seeding completed")
	} else {
		fmt.Println("No operation specified")
	}

	// ルーターの設定
	r := router.SetupRouter(db)

	r.Run(":8080")
}
