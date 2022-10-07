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
	if Order.Id != 0 {
		var existingItem []model.Item
		c.connection.Where("order_id = ?", Order.Id).Find(&existingItem)
		var existingItemCount int = len(existingItem)
		var newItemCount int = len(newOrder.Items)
		if newItemCount < existingItemCount {
			c.connection.Debug().Unscoped().Model(model.Item{}).Where("order_id = ?", Order.Id).
			Order("id asc").Limit(existingItemCount - newItemCount).Offset(newItemCount).Delete(&model.Item{})
		}
	}
	Order.CustomerName = newOrder.CustomerName
	err := c.connection.Save(Order).Error
	if err != nil {
		return err
	}

	for eachIndex, eachItem := range newOrder.Items {
		var item model.Item
		c.connection.Where("order_id = ?", Order.Id).Order("id asc").Limit(1).Offset(eachIndex).Find(&item)
		item.OrderID = Order.Id
		item.ItemCode = eachItem.ItemCode
		item.Description = eachItem.Description
		item.Quantity = eachItem.Quantity
		err := c.connection.Save(&item).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (c ConnectionDB) getOrderDetailById(id uint, preload bool) (model.Order, error) {
	var order model.Order
	conn := c.connection.Where("id = ?", id)
	if preload {
		conn = conn.Preload("Items")
	}
	conn = conn.First(&order)
	err := conn.Error
	if err != nil {
		return order, err
	}
	return order, err

}

func (c ConnectionDB) deleteOrder(id int) error {
	var order model.Order
	err := c.connection.Delete(&order, id).Error
	if err != nil {
		return err
	}
	return nil
}




