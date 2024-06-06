package data

import (
	doctor "PetPalApp/features/doctor/data"
	product "PetPalApp/features/product/data"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FullName       string
	Email          string 				`gorm:"unique"`
	Password       string
	NumberPhone    *string 				`gorm:"unique"`
	Address        *string
	ProfilePicture string 
	Coordinate     *string
	Role           string 				`gorm:"default:'admin'"`
	Products	   []product.Product 	`gorm:"foreign_key:AdminID"`
	Doctors		   []doctor.Doctor 		`gorm:"foreign_key:AdminID"`
}
