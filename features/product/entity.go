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
	SelectAll() ([]Core, error)
	SelectAllAdmin(userid uint) ([]Core, error)
	SelectById(id uint) (*Core, error)
	SelectByIdAdmin(id uint, userid uint) (*Core, error)
	PutById(id uint, userid uint, input Core) error
}

type ServiceInterface interface {
	Create(id uint, input Core, file io.Reader, handlerFilename string) (string, error)
	GetAll() ([]Core, error)
	GetAllAdmin(userid uint) ([]Core, error)
	GetProductById(id uint) (data *Core, err error)
	GetProductByIdAdmin(id uint, userid uint) (data *Core, err error)
	UpdateById(id uint, userid uint, input Core) error
}
