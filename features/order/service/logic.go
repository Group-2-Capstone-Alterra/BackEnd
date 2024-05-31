package service

import (
	"PetPalApp/features/order"
)

type OrderService struct {
    orderModel order.OrderModel
}

func New(om order.OrderModel) order.OrderService {
    return &OrderService{
        orderModel: om,
    }
}

func (os *OrderService) CreateOrder(order order.OrderCore) error {
    return os.orderModel.CreateOrder(order)
}