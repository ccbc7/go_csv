package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"project/config"
	"project/internal/database/seeders"
	"project/internal/ent"
	"project/internal/router"

	_ "github.com/lib/pq"
)

// @title			Swagger Example API
// @version		1.0
// @description	This is a sample server for a pet store.
// @BasePath		/api/v1
func main() {
	seed := flag.Bool("seed", false, "Run the database seeders")
	migrate := flag.Bool("migrate", false, "Run ent schema migration and exit")
	flag.Parse()

	// 初期化(環境変数の読み込み)
	config.Initialize()

	// Entクライアントの作成
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if *migrate {
		// Entのマイグレーションのみ実行して終了
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		fmt.Println("Migration completed")
		return
	}

	if *seed {
		seeders.SeedAll(client)
		fmt.Println("Seeding completed")
		return
	}

	// ルーターの設定
	r := router.SetupRouter(client)

	r.Run(":8080")
}

// configEnv は環境変数を読む簡易ヘルパー
func configEnv(key string) string {
	return getenv(key)
}

func getenv(key string) string {
	if v := stringLookupEnv(key); v != "" {
		return v
	}
	return ""
}

func stringLookupEnv(key string) string {
	// os.LookupEnv を直接使うと import が増えるため、ここで薄いラッパを定義
	// 既存ファイルの import を大きく変えないようにする
	return getEnvFromOS(key)
}

// 実体は別関数に分ける（テスト容易性確保用の分岐を残せる）
func getEnvFromOS(key string) string {
	return func(k string) string { v, _ := lookupEnv(k); return v }(key)
}

// 標準ライブラリの os.LookupEnv に委譲
func lookupEnv(key string) (string, bool) {
	return osLookupEnv(key)
}

// ここでのみ os を使用（import 追加を局所化）
func osLookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
