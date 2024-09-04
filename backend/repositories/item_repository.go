package repositories

import (
	"errors"

	"project/models"

	"gorm.io/gorm"
)

// インターフェースを定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updatedItem models.Item) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

// 構造体を定義 構造体とは、フィールドの集まりを定義した型
type ItemMemoryRepository struct {
	items []models.Item
}

// コンストラクタを定義
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// メソッドを定義 このメソッドは、ItemMemoryRepository構造体のItemsフィールドを返す
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

// 構造体のItemsフィールドからIDを検索して返す
func (r *ItemMemoryRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

// 更新
func (r *ItemMemoryRepository) Update(updatedItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updatedItem.ID {
			r.items[i] = updatedItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

// 削除
func (r *ItemMemoryRepository) Delete(itemId uint, userId uint) error {
	// i, v := range r.items でスライスのインデックスと要素を取得
	for i, v := range r.items {
		if v.ID == itemId {
			// append()関数でスライスの要素を削除 ...演算子でスライスの要素を展開
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

/*
* DBを使う場合のリポジトリの実装
 */

// 構造体を定義
type ItemDBRepository struct {
	db *gorm.DB
}

// 作成
func (r *ItemDBRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// 削除
func (r *ItemDBRepository) Delete(itemId uint, userId uint) error {
	deleteItem, err := r.FindById(itemId, userId)
	if err != nil {
		return err
	}

	result := r.db.Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 全件取得
func (r *ItemDBRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// IDで取得
func (r *ItemDBRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	var item models.Item
	// 第一引数に構造体のポインタ、第二引数にIDを指定
	result := r.db.First(&item, "id = ? AND user_id = ?", itemId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

// 更新
func (r *ItemDBRepository) Update(updatedItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updatedItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedItem, nil
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemDBRepository{db: db}
}
