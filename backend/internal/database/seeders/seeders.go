package seeders

import (
	"context"
	"log"

	"project/internal/ent"
)

func SeedAll(client *ent.Client) {
	// 一度データを削除（外部キー制約のため、アイテムを先に削除）
	_, err := client.Item.Delete().Exec(context.Background())
	if err != nil {
		log.Fatalf("アイテムの削除に失敗しました: %v", err)
	}
	log.Println("アイテムの削除が完了しました")

	_, err = client.User.Delete().Exec(context.Background())
	if err != nil {
		log.Fatalf("ユーザーの削除に失敗しました: %v", err)
	}
	log.Println("ユーザーの削除が完了しました")

	seeds := []func(*ent.Client) error{
		SeedUsers,
		SeedItems,
	}

	for _, seed := range seeds {
		if err := seed(client); err != nil {
			log.Fatalf("failed to seed: %v", err)
		}
	}

	log.Println("All seeding completed successfully")
}
