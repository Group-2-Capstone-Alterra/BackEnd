package product

import "time"

type Core struct {
	ID          uint
	IdUser      uint
	ProductName string
	Price       float32
	Stock       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
