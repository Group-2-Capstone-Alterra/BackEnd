package user

import (
	"io"
	"time"
)

type Core struct {
	ID             uint
	FullName       string
	Email          string
	NumberPhone    string
	Address        string
	Password       string
	ProfilePicture string
	Coordinate     string
	Token          string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectByEmail(email string) (*Core, error)
	SelectById(id uint) (*Core, error)
	PutById(id uint, input Core) error
	Delete(id uint) error
}

type ServiceInterface interface {
	Create(input Core) error
	Login(email string, password string) (data *Core, token string, err error)
	GetProfile(id uint) (data *Core, err error)
	UpdateById(id uint, input Core, file io.Reader, handlerFilename string) (string, error)
	Delete(id uint) error
}
