package service

import (
	order "PetPalApp/features/order"
)

type OrderService struct {
	OrderModel order.OrderModel
}

func New(om order.OrderModel) order.OrderService {
	return &OrderService{
		OrderModel: om,
	}
}

func (os *OrderService) CreateOrder(opCore order.Order) (order.Order, error) {
    return os.OrderModel.CreateOrder(opCore)
}


func (os *OrderService) GetOrdersByUserID(userID uint) ([]order.Order, error) {
    return os.OrderModel.GetOrdersByUserID(userID)
}

func (os *OrderService) GetProductByID(id uint) (data *order.Product, err error) {	
	return os.OrderModel.GetProductByID(id)	
}

func (os *OrderService) GetOrderByID(orderID uint) (*order.Order, error) {
    return os.OrderModel.GetOrderByID(orderID)
}
