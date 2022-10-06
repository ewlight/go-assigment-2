package app

import (
	"gorm.io/gorm"
	"model"
	"resources"
)

type OrderService interface {
	getOrderList() ([]model.Order, error, int64)
	addOrder(Order *model.Order, newOrder resources.InputOrder) error
	getOrderDetailById(id uint, preload bool) (model.Order, error)
	deleteOrder(id int) error
}

type ConnectionDB struct {
	connection *gorm.DB
}

func newOrderService() OrderService  {
	return &ConnectionDB{connection: dbConnection()}
}

func (c ConnectionDB) getOrderList() ([]model.Order, error, int64) {
	var order []model.Order
	var count int64
	connection := c.connection.Model(&order).Preload("Items").Find(&order)
	err := connection.Error
	if err != nil {
		return order, err, 0
	}
	c.connection.Model(order).Count(&count)
	return order, nil, count
}

func (c ConnectionDB) addOrder(Order *model.Order, newOrder resources.InputOrder) error {
	panic("implement me")
}

func (c ConnectionDB) getOrderDetailById(id uint, preload bool) (model.Order, error) {
	panic("implement me")
}

func (c ConnectionDB) deleteOrder(id int) error {
	panic("implement me")
}




