package seeders

import (
	"context"

	"project/internal/ent"
)

func SeedItems(client *ent.Client) error {
	// 既存のユーザーを取得
	users, err := client.User.Query().All(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return nil // ユーザーがいない場合は何もしない
	}

	items := []struct {
		Name        string
		Price       int
		Description string
		SoldOut     bool
	}{
		{Name: "Item1", Price: 100, Description: "Description1", SoldOut: false},
		{Name: "Item2", Price: 200, Description: "Description2", SoldOut: false},
		{Name: "Item3", Price: 300, Description: "Description3", SoldOut: false},
	}

	for i, item := range items {
		userID := users[i%len(users)].ID // ユーザーを循環して使用
		_, err := client.Item.Create().
			SetName(item.Name).
			SetPrice(item.Price).
			SetDescription(item.Description).
			SetSoldOut(item.SoldOut).
			SetUserID(userID).
			Save(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}
