package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primaryKey;auto_increment;column:id"`
	CustomerName string    `json:"customer_name" gorm:"column:customer_name"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"column:ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UpdateOrder struct {
	CustomerName string       `json:"customer_name" gorm:"column:customer_name"`
	OrderedAt    time.Time    `json:"ordered_at" gorm:"column:ordered_at"`
	Items        []UpdateItem `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
