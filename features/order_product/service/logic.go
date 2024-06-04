package service

import (
	"PetPalApp/features/order_product"
)

type OrderProductService struct {
	OrderProductModel order_product.OrderProductModel
}

func New(opm order_product.OrderProductModel) order_product.OrderProductService {
	return &OrderProductService{
		OrderProductModel: opm,
	}
}

func (ops *OrderProductService) CreateOrderProduct(opCore order_product.OrderProductCore) error {
	return ops.OrderProductModel.CreateOrderProduct(opCore)
}
