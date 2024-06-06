package product

import (
	"io"
	"time"
)

type Core struct {
	ID             uint
	AdminID        uint
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
	SelectAll(offset uint, sortStr string) ([]Core, error)
	SelectAllAdmin(adminid uint, offset uint) ([]Core, error)
	SelectById(id uint) (*Core, error)
	SelectByIdAdmin(id uint, adminid uint) (*Core, error)
	PutById(id uint, adminid uint, input Core) error
	Delete(id uint, adminid uint) error
	VerIsAdmin(adminid uint) (*Core, error)
}

type ServiceInterface interface {
	Create(id uint, input Core, file io.Reader, handlerFilename string) (string, error)
	GetAll(adminid uint, role string, offset uint, sortStr string) ([]Core, error)
	GetProductById(id uint, adminid uint) (data *Core, err error)
	UpdateById(id uint, adminid uint, input Core, file io.Reader, handlerFilename string) (string, error)
	Delete(id uint, adminid uint) error
}
