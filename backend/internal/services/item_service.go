package services

import (
	"project/internal/dto"
	"project/internal/ent"
	"project/internal/repositories"
)

// インターフェースを定義
type ItemService interface {
	FindAll() ([]*ent.Item, error)
	FindById(itemId int, userId int) (*ent.Item, error)
	Create(createItemInput dto.CreateItemInput, userId int) (*ent.Item, error)
	Update(itemId int, updateItemInput dto.UpdateItemInput, userId int) (*ent.Item, error)
	Delete(itemId int, userId int) error
}

// 構造体を定義
type itemService struct {
	repository repositories.ItemRepository
}

// コンストラクタを定義
func NewItemService(repository repositories.ItemRepository) ItemService {
	return &itemService{repository: repository}
}

// 全ての商品を取得
func (s *itemService) FindAll() ([]*ent.Item, error) {
	return s.repository.FindAll()
}

// IDで商品を取得
func (s *itemService) FindById(itemId int, userId int) (*ent.Item, error) {
	return s.repository.FindById(itemId, userId)
}

// 作成
func (s *itemService) Create(createItemInput dto.CreateItemInput, userId int) (*ent.Item, error) {
	// リポジトリ層のCreateメソッドを呼び出し、作成処理を行う
	return s.repository.Create(createItemInput.Name, int(createItemInput.Price), createItemInput.Description, false, userId)
}

// 更新
func (s *itemService) Update(itemId int, updateItemInput dto.UpdateItemInput, userId int) (*ent.Item, error) {
	// IDとユーザーIDで商品を取得,ユーザは自分の商品のみ更新できる
	targetItem, err := s.FindById(itemId, userId)
	if err != nil {
		return nil, err
	}

	name := targetItem.Name
	price := targetItem.Price
	description := targetItem.Description
	soldOut := targetItem.SoldOut

	if updateItemInput.Name != nil {
		name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		price = int(*updateItemInput.Price)
	}
	if updateItemInput.Description != nil {
		description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		soldOut = *updateItemInput.SoldOut
	}

	// リポジトリ層のUpdateメソッドを呼び出し、更新処理を行う
	return s.repository.Update(itemId, name, price, description, soldOut, userId)
}

// 削除
func (s *itemService) Delete(itemId int, userId int) error {
	return s.repository.Delete(itemId, userId)
}
