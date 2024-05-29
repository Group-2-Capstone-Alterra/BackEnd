package data

import (
	"PetPalApp/features/user"
	"PetPalApp/utils/helper"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       *string `json:"FullName" form:"FullName"`
	Email          *string `gorm:"unique" json:"email" form:"email"`
	NumberPhone    *string `gorm:"unique" json:"NumberPhone" form:"NumberPhone"`
	Address        *string `json:"Address" form:"Address"`
	Password       *string `json:"Password" form:"Password"`
	ProfilePicture *string `json:"ProfilePicture" form:"ProfilePicture"`
}

func UserCoreToUserGorm(userCore user.Core, helper helper.HelperInterface) User {
	userGorm := User{
		FullName:       helper.ConvertToNullableString(userCore.FullName),
		Email:          helper.ConvertToNullableString(userCore.Email),
		NumberPhone:    helper.ConvertToNullableString(userCore.NumberPhone),
		Address:        helper.ConvertToNullableString(userCore.Address),
		Password:       helper.ConvertToNullableString(userCore.Password),
		ProfilePicture: helper.ConvertToNullableString(userCore.ProfilePicture),
	}
	return userGorm
}

func UserGormToUserCore(userGorm User, helper helper.HelperInterface) user.Core {
	userCore := user.Core{
		ID:             userGorm.ID,
		FullName:       helper.DereferenceString(userGorm.FullName),
		Email:          helper.DereferenceString(userGorm.Email),
		NumberPhone:    helper.DereferenceString(userGorm.NumberPhone),
		Address:        helper.DereferenceString(userGorm.Address),
		Password:       helper.DereferenceString(userGorm.Password),
		ProfilePicture: helper.DereferenceString(userGorm.ProfilePicture),
		CreatedAt:      userGorm.CreatedAt,
		UpdatedAt:      userGorm.UpdatedAt,
	}
	return userCore
}
