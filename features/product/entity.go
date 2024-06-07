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
	Distance       float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll(limit, offset uint, sortStr string) ([]Core, error)
	SelectAllAdmin(limit, userid uint, offset uint) ([]Core, error)
	SelectById(id uint) (*Core, error)
	SelectByIdAdmin(id uint, userid uint) (*Core, error)
	PutById(id uint, userid uint, input Core) error
	Delete(id uint, userid uint) error
	VerIsAdmin(userid uint) (*Core, error)
}

type ServiceInterface interface {
	Create(id uint, input Core, file io.Reader, handlerFilename string) (string, error)
	GetAll(userid uint, limit uint, role string, offset uint, sortStr string) ([]Core, error)
	GetProductById(id uint, userid uint) (data *Core, err error)
	UpdateById(id uint, userid uint, input Core, file io.Reader, handlerFilename string) (string, error)
	Delete(id uint, userid uint) error
}
