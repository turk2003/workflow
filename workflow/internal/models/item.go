package models

import (
	constants "github.com/turk2003/workflow/internal/constant"
)

type Item struct {
	ID       uint                 `gorm:"primaryKey" json:"id"`
	Title    string               `json:"title"`
	Amount   int                  `json:"amount"`
	Quantity int                  `json:"quantity"`
	Status   constants.ItemStatus `json:"status"`
	OwnerID  int                  `gorm:"primaryKey" json:"owner_id"`
}
