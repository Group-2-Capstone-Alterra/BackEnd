package data

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FullName       string
	Email          string `gorm:"unique"`
	Password       string
	NumberPhone    *string `gorm:"unique"`
	Address        *string
	ProfilePicture string 
	Coordinate     *string
	Role           string `gorm:"default:'admin'"`
}
