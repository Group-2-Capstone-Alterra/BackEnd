package order

import (
	_productHandler "PetPalApp/features/product/handler"
	_userHandler "PetPalApp/features/user/handler"

	"time"
)

type OrderCore struct {
	ID        uint
	UserID    uint
	ProductID uint
	Quantity  int
	Total     float64
	Status    string
	User      _userHandler.OrderResponse
	Product   _productHandler.OrderResponse

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
