package repository

import (
	"errors"
	"hactive/assigment-2/entity"
	"hactive/assigment-2/interfaces"
	"sync"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateOrder(order *entity.Order) error {
	result := r.DB.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) GetOrders() ([]entity.Order, error) {
	orders := []entity.Order{}

	result := r.DB.Preload("Items").Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (r *repository) UpdateOrder(id uint, order *entity.Order) error {
	orders := entity.Order{}
	item := entity.Item{}
	wg := sync.WaitGroup{}

	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		// update order
		result := r.DB.Model(&orders).Where("id = ?", id).Updates(order)
		if result.RowsAffected == 0 {
			err = errors.New("order not found")
		}
	}()

	go func() {
		defer wg.Done()
		// update order items
		for _, v := range order.Items {
			result := r.DB.Model(&item).Where("order_id = ?", id).Updates(v)
			if result.RowsAffected == 0 {
				err = errors.New("order item not found")
			}
		}
	}()

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteOrder(id uint) error {
	order := entity.Order{}
	item := entity.Item{}
	wg := sync.WaitGroup{}

	var err error

	wg.Add(2)
	go func() {
		defer wg.Done()
		// delete order
		result := r.DB.Delete(&order, id)
		if result.RowsAffected == 0 {
			err = errors.New("order not found")
		}
	}()

	go func() {
		defer wg.Done()
		// delete order items
		result := r.DB.Delete(&item, "order_id = ?", id)
		if result.RowsAffected == 0 {
			err = errors.New("order items not found")
		}
	}()

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}
