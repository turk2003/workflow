package services

import (
	constants "github.com/turk2003/workflow/internal/constant"
	"github.com/turk2003/workflow/internal/models"
	repositories "github.com/turk2003/workflow/internal/respositories"
)

type ItemService interface {
	CreateItem(item models.Item) error
	GetAllItems() ([]models.Item, error)
	GetItemByID(id uint) (models.Item, error)
	UpdateItem(item models.Item) error
	DeleteItem(id uint) error
}

type itemService struct {
	repository repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repository: repo}
}

func (s *itemService) CreateItem(item models.Item) error {
	// Set default status to PENDING if not provided
	if item.Status == "" {
		item.Status = constants.ItemPendingStatus
	}
	return s.repository.CreateItem(item)
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repository.GetAllItems()
}

func (s *itemService) GetItemByID(id uint) (models.Item, error) {
	return s.repository.GetItemByID(id)
}

func (s *itemService) UpdateItem(item models.Item) error {
	return s.repository.UpdateItem(item)
}

func (s *itemService) DeleteItem(id uint) error {
	return s.repository.DeleteItem(id)
}
