package service

import (
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

func (ps *PaymentService) CreatePayment(payment payment.PaymentCore) error {
	return ps.paymentModel.Create(payment)
}
