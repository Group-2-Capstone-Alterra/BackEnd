package order

import (
	"time"
)

type OrderCore struct {
	ID        uint
	UserID    uint
	Total     float64
	Status    string
	CreatedAt time.Time
}

type OrderModel interface {
	CreateOrder(OrderCore) error
	GetOrdersByUserID(userID uint) ([]OrderCore, error)
}

type OrderService interface {
	CreateOrder(OrderCore) error
	GetOrdersByUserID(userID uint) ([]OrderCore, error)
}
