package data

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	FullName       string
	Email          string
	Specialization string
}