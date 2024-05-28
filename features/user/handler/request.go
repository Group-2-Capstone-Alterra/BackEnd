package handler

import (
	"PetPalApp/features/user"

	"gorm.io/gorm"
)

type UserRequest struct {
	gorm.Model
	FullName       string `json:"FullName" form:"FullName"`
	Email          string `gorm:"unique" json:"email" form:"email"`
	NumberPhone    string `gorm:"unique" json:"NumberPhone" form:"NumberPhone"`
	Address        string `json:"Address" form:"Address"`
	Password       string `json:"Password" form:"Password"`
	RetypePassword string `json:"RetypePassword" form:"RetypePassword"`
	ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
}
type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	inputCore := user.Core{
		FullName:       input.FullName,
		Email:          input.Email,
		Password:       input.Password,
		RetypePassword: input.RetypePassword,
	}
	return inputCore
}
