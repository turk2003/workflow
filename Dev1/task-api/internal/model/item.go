// internal/model/item.go

package model

import "task-api/internal/constant"

type Item struct {
	ID       uint                `gorm:"primaryKey" json:"id"`
	Title    string              `json:"title"`
	Price    float64             `json:"price"`
	Quantity uint                `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
}
