package product

import (
	"io"
	"time"
)

type Core struct {
	ID             uint
	IdUser         uint
	ProductName    string
	Price          float32
	Stock          uint
	Description    string
	ProductPicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DataInterface interface {
	Insert(input Core) error
}

type ServiceInterface interface {
	Create(id uint, input Core, file io.Reader, handlerFilename string) (string, error)
}
