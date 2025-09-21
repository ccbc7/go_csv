package seeders

import (
	"context"
	"fmt"
	"log"

	"project/internal/ent"
)

// シーダー関数マップ
var seeders = map[string]func(*ent.Client) error{
	"users":  SeedUsers,
	"items":  SeedItems,
	"movies": SeedMovies,
}

// Delete順序（外部キー制約を考慮）
var deleteOrder = []string{"items", "users", "movies"}

// SeedAll はシードを実行
func SeedAll(client *ent.Client) {
	log.Println("シード処理を開始します...")

	// 既存データを削除
	log.Println("データ削除処理を開始します...")
	for _, table := range deleteOrder {
		log.Printf("テーブルを削除中: %s", table)
		if err := deleteTable(client, table); err != nil {
			log.Fatalf("%s の削除に失敗しました: %v", table, err)
		}
		log.Printf("%s の削除が完了しました", table)
	}

	// シード実行
	log.Println("シード実行処理を開始します...")
	for name, seeder := range seeders {
		log.Printf("%s をシード中...", name)
		if err := seeder(client); err != nil {
			log.Fatalf("%s のシードに失敗しました: %v", name, err)
		}
		log.Printf("%s のシードが完了しました", name)
	}

	log.Println("全てのシード処理が正常に完了しました")
}

// deleteTable はテーブルごとの削除を安全に実行
func deleteTable(client *ent.Client, table string) error {
	switch table {
	case "users":
		count, err := client.User.Delete().Exec(context.Background())
		if err != nil {
			return err
		}
		log.Printf("%d 件のユーザーを削除しました", count)
		return nil
	case "items":
		count, err := client.Item.Delete().Exec(context.Background())
		if err != nil {
			return err
		}
		log.Printf("%d 件のアイテムを削除しました", count)
		return nil
	case "movies":
		count, err := client.Movie.Delete().Exec(context.Background())
		if err != nil {
			return err
		}
		log.Printf("%d 件の映画を削除しました", count)
		return nil
	default:
		return fmt.Errorf("不明なテーブル: %s", table)
	}
}
