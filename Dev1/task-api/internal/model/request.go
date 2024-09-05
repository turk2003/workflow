// internal/model/request.go
package model

import "task-api/internal/constant"

type RequestItem struct {
	Title    string
	Price    float64
	Quantity uint
}

type RequestFindItem struct {
	Statuses []constant.ItemStatus `form:"status[]"`
}

// internal/model/request.go

type RequestUpdateItem struct {
	Status constant.ItemStatus
}
