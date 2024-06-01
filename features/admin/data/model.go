package data

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FullName       string
	Email          string `gorm:"unique"`
	NumberPhone    string `gorm:"unique"`
	Address        string
	Password       string
	ProfilePicture string
	Coordinate     string
}
