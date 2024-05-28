package data

import (
	"PetPalApp/features/user"
	"PetPalApp/utils/helper"

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

type UserRegister struct {
	gorm.Model
	FullName       *string `json:"FullName" form:"FullName"`
	Email          *string `gorm:"unique" json:"email" form:"email"`
	NumberPhone    *string `gorm:"unique" json:"NumberPhone" form:"NumberPhone"`
	Address        *string `json:"Address" form:"Address"`
	Password       *string `json:"Password" form:"Password"`
	ProfilePicture *string `json:"ProfilePicture" form:"ProfilePicture"`
}

func UserCoreToUserGorm(userCore user.Core, helper helper.HelperInterface) UserRegister {
	userGorm := UserRegister{
		FullName:       helper.ConvertToNullableString(userCore.FullName),
		Email:          helper.ConvertToNullableString(userCore.Email),
		NumberPhone:    helper.ConvertToNullableString(userCore.NumberPhone),
		Address:        helper.ConvertToNullableString(userCore.Address),
		Password:       helper.ConvertToNullableString(userCore.Password),
		ProfilePicture: helper.ConvertToNullableString(userCore.ProfilePicture),
	}
	return userGorm
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
