package user

import (
	"time"
)

type Core struct {
	ID             uint
	FullName       string
	Email          string
	NumberPhone    string
	Address        string
	Password       string
	RetypePassword string
	ProfilePicture string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DataInterface interface {
	Insert(input Core) error
}

type ServiceInterface interface {
	Create(input Core) (string, error)
}
