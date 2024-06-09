package service

import (
	"PetPalApp/features/order"
	"PetPalApp/features/payment"
)

type PaymentService struct {
	paymentModel payment.PaymentModel
}

type OrderService struct {
	OrderModel order.OrderModel
}

func New(pm payment.PaymentModel) payment.PaymentService {
	return &PaymentService{
		paymentModel: pm,
	}
}

func (ps *PaymentService) CreatePayment(payment payment.PaymentCore) error {
	return ps.paymentModel.Create(payment)
}

func (ps *PaymentService) GetOrderById(id uint) (data *payment.Order, err error) {	
	return ps.paymentModel.GetOrderByID(id)	
}

func (ps *PaymentService) GetUserById(id uint) (data *payment.User, err error) {	
	return ps.paymentModel.GetUserByID(id)	
}
