package repositories

import (
	"github.com/turk2003/workflow/internal/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item models.Item) error
	GetAllItems() ([]models.Item, error)
	GetItemByID(id uint) (models.Item, error)
	UpdateItem(item models.Item) error
	DeleteItem(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) CreateItem(item models.Item) error {
	return r.db.Create(&item).Error
}

func (r *itemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) GetItemByID(id uint) (models.Item, error) {
	var item models.Item
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *itemRepository) UpdateItem(item models.Item) error {
	return r.db.Save(&item).Error
}

func (r *itemRepository) DeleteItem(id uint) error {
	return r.db.Delete(&models.Item{}, id).Error
}
