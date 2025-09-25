package repositories

import (
	"context"

	"project/internal/ent"
	"project/internal/ent/item"
)

// インターフェースを定義
type ItemRepository interface {
	FindAll() ([]*ent.Item, error)
	FindById(itemId int, userId int) (*ent.Item, error)
	Create(name string, price int, description string, soldOut bool, userId int) (*ent.Item, error)
	Update(itemId int, name string, price int, description string, soldOut bool, userId int) (*ent.Item, error)
	Delete(itemId int, userId int) error
}

// Entクライアントを使用するリポジトリ
type itemEntRepository struct {
	client *ent.Client
}

func (r *itemEntRepository) FindAll() ([]*ent.Item, error) {
	return r.client.Item.Query().All(context.Background())
}

func (r *itemEntRepository) FindById(itemId int, userId int) (*ent.Item, error) {
	return r.client.Item.Query().
		Where(item.And(item.IDEQ(itemId), item.UserIDEQ(userId))).
		Only(context.Background())
}

func (r *itemEntRepository) Create(name string, price int, description string, soldOut bool, userId int) (*ent.Item, error) {
	return r.client.Item.Create().
		SetName(name).
		SetPrice(price).
		SetDescription(description).
		SetSoldOut(soldOut).
		SetUserID(userId).
		Save(context.Background())
}

func (r *itemEntRepository) Update(itemId int, name string, price int, description string, soldOut bool, userId int) (*ent.Item, error) {
	return r.client.Item.UpdateOneID(itemId).
		SetName(name).
		SetPrice(price).
		SetDescription(description).
		SetSoldOut(soldOut).
		SetUserID(userId).
		Save(context.Background())
}

func (r *itemEntRepository) Delete(itemId int, userId int) error {
	return r.client.Item.DeleteOneID(itemId).Exec(context.Background())
}

func NewItemRepository(client *ent.Client) ItemRepository {
	return &itemEntRepository{client: client}
}
