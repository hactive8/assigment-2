package entity

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey;auto_increment;column:id"`
	ItemCode    string `json:"item_code" gorm:"column:item_code"`
	Description string `json:"description" gorm:"column:description"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
	OrderID     int    `json:"order_id" gorm:"column:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateItem struct {
	ItemCode    string `json:"item_code" gorm:"column:item_code"`
	Description string `json:"description" gorm:"column:description"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
	OrderID     int    `json:"order_id" gorm:"column:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
