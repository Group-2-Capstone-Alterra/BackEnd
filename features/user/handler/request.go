package handler

import (
	"PetPalApp/features/user"

	"gorm.io/gorm"
)

type UserRequest struct {
	gorm.Model
	FullName       string `json:"full_name" form:"full_name"`
	Email          string `gorm:"unique" json:"email" form:"email"`
	NumberPhone    string `gorm:"unique" json:"number_phone" form:"number_phone"`
	Address        string `json:"address" form:"address"`
	Password       string `json:"password" form:"password"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	inputCore := user.Core{
		FullName:       input.FullName,
		Email:          input.Email,
		NumberPhone:    input.NumberPhone,
		Address:        input.Address,
		Password:       input.Password,
		ProfilePicture: input.ProfilePicture,
	}
	return inputCore
}
