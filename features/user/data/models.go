package data

import (
	"PetPalApp/features/user"

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

func UserGormToUserCore(userGorm User) user.Core {
	userCore := user.Core{
		ID:             userGorm.ID,
		FullName:       userGorm.FullName,
		Email:          userGorm.Email,
		NumberPhone:    userGorm.NumberPhone,
		Address:        userGorm.Address,
		Password:       userGorm.Password,
		ProfilePicture: userGorm.ProfilePicture,
		CreatedAt:      userGorm.CreatedAt,
		UpdatedAt:      userGorm.UpdatedAt,
	}
	return userCore
}

func UserCoreToUserGorm(userCore user.Core) User {
	userGorm := User{
		FullName:       userCore.FullName,
		Email:          userCore.Email,
		Password:       userCore.Password,
		ProfilePicture: userCore.ProfilePicture,
	}
	return userGorm
}
