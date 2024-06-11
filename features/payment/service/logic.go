package service

import (
	"PetPalApp/features/order"
	"PetPalApp/features/payment"
)

type PaymentService struct {
	paymentModel payment.PaymentModel
}

func New(pm payment.PaymentModel) payment.PaymentService {
	return &PaymentService{
		paymentModel: pm,
	}
}

func (ps *PaymentService) FindOrCreatePayment(orderID uint, payment payment.Payment) (payment.Payment, error) {
	return ps.paymentModel.FindOrCreatePayment(orderID, payment)
}

func (ps *PaymentService) GetPaymentByID(id uint) (data *payment.Payment, err error) {	
	return ps.paymentModel.GetPaymentByID(id)	
}

func (ps *PaymentService) GetOrderByID(id uint) (data *order.Order, err error) {	
	return ps.paymentModel.GetOrderByID(id)	
}

func (ps *PaymentService) GetUserByID(id uint) (data *payment.User, err error) {	
	return ps.paymentModel.GetUserByID(id)	
}

func (ps *PaymentService) Update(orderID uint, payment payment.Payment) (payment.Payment, error) {	
	return ps.paymentModel.Update(orderID, payment)	
}

