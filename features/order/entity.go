package order

import "time"

type OrderCore struct {
    ID        uint
    UserID    uint
    ProductID uint
    Quantity  int
    Total     float64
    Status    string
    CreatedAt time.Time
}

type OrderModel interface {
    CreateOrder(OrderCore) error
}

type OrderService interface {
    CreateOrder(OrderCore) error
}
