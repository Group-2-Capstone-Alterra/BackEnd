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

func (ps *PaymentService) CreatePayment(payment payment.Payment) (payment.Payment, error) {
	return ps.paymentModel.Create(payment)
}

func (ps *PaymentService) GetPaymentByID(id uint) (data *payment.Payment, err error) {	
	return ps.paymentModel.GetPaymentByID(id)	
}

func (ps *PaymentService) GetOrderByID(id uint) (data *payment.Order, err error) {	
	return ps.paymentModel.GetOrderByID(id)	
}

func (ps *PaymentService) GetUserByID(id uint) (data *payment.User, err error) {	
	return ps.paymentModel.GetUserByID(id)	
}

