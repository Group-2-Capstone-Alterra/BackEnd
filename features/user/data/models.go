package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `json:"FullName" form:"FullName"`
	Email          string `gorm:"unique" json:"email" form:"email"`
	NumberPhone    string `gorm:"unique" json:"NumberPhone" form:"NumberPhone"`
	Address        string `json:"Address" form:"Address"`
	Password       string `json:"Password" form:"Password"`
	ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
}
