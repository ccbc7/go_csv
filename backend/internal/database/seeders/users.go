package seeders

import (
	"context"

	"project/internal/ent"
)

func SeedUsers(client *ent.Client) error {
	users := []struct {
		Name     string
		LoginID  string
		Password string
	}{
		{Name: "テストユーザー1", LoginID: "test1@example.com", Password: "password123"},
		{Name: "テストユーザー2", LoginID: "test2@example.com", Password: "password123"},
	}

	for _, user := range users {
		_, err := client.User.Create().
			SetName(user.Name).
			SetLoginID(user.LoginID).
			SetPassword(user.Password).
			Save(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}
