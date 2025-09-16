package repositories

import (
	"context"

	"project/internal/ent"
	"project/internal/ent/item"
)

// インターフェースを定義
type IItemRepository interface {
	FindAll() ([]*ent.Item, error)
	FindById(itemId int, userId int) (*ent.Item, error)
	Create(name string, price int, description string, soldOut bool, userId int) (*ent.Item, error)
	Update(itemId int, name string, price int, description string, soldOut bool, userId int) (*ent.Item, error)
	Delete(itemId int, userId int) error
}

// Entクライアントを使用するリポジトリ
type ItemEntRepository struct {
	client *ent.Client
}

func (r *ItemEntRepository) FindAll() ([]*ent.Item, error) {
	return r.client.Item.Query().All(context.Background())
}

func (r *ItemEntRepository) FindById(itemId int, userId int) (*ent.Item, error) {
	return r.client.Item.Query().
		Where(item.And(item.IDEQ(itemId), item.UserIDEQ(userId))).
		Only(context.Background())
}

func (r *ItemEntRepository) Create(name string, price int, description string, soldOut bool, userId int) (*ent.Item, error) {
	return r.client.Item.Create().
		SetName(name).
		SetPrice(price).
		SetDescription(description).
		SetSoldOut(soldOut).
		SetUserID(userId).
		Save(context.Background())
}

func (r *ItemEntRepository) Update(itemId int, name string, price int, description string, soldOut bool, userId int) (*ent.Item, error) {
	return r.client.Item.UpdateOneID(itemId).
		SetName(name).
		SetPrice(price).
		SetDescription(description).
		SetSoldOut(soldOut).
		SetUserID(userId).
		Save(context.Background())
}

func (r *ItemEntRepository) Delete(itemId int, userId int) error {
	return r.client.Item.DeleteOneID(itemId).Exec(context.Background())
}

func NewItemRepository(client *ent.Client) IItemRepository {
	return &ItemEntRepository{client: client}
}
