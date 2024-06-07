package data

import (
	order "PetPalApp/features/order/data"
	transaction "PetPalApp/features/transaction/data"
	"PetPalApp/features/user"
	"PetPalApp/utils/helperuser"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       *string
	Email          *string 						`gorm:"unique"`
	NumberPhone    *string 						`gorm:"unique"`
	Address        *string
	Password       *string
	ProfilePicture *string
	Coordinate     *string
	Role           string 						`gorm:"default:'user'"`
	Orders		   []order.Order 				`gorm:"foreign_key:UserID"`
	Transactions   []transaction.Transaction 	`gorm:"foreign_key:UserID"`
}

func UserCoreToUserGorm(userCore user.Core, helper helperuser.HelperuserInterface) User {
	userGorm := User{
		FullName:       helper.ConvertToNullableString(userCore.FullName),
		Email:          helper.ConvertToNullableString(userCore.Email),
		NumberPhone:    helper.ConvertToNullableString(userCore.NumberPhone),
		Address:        helper.ConvertToNullableString(userCore.Address),
		Password:       helper.ConvertToNullableString(userCore.Password),
		ProfilePicture: helper.ConvertToNullableString(userCore.ProfilePicture),
		Coordinate:     helper.ConvertToNullableString(userCore.Coordinate),
	}
	return userGorm
}

func UserGormToUserCore(userGorm User, helper helperuser.HelperuserInterface) user.Core {
	userCore := user.Core{
		ID:             userGorm.ID,
		FullName:       helper.DereferenceString(userGorm.FullName),
		Email:          helper.DereferenceString(userGorm.Email),
		NumberPhone:    helper.DereferenceString(userGorm.NumberPhone),
		Address:        helper.DereferenceString(userGorm.Address),
		Password:       helper.DereferenceString(userGorm.Password),
		ProfilePicture: helper.DereferenceString(userGorm.ProfilePicture),
		Coordinate:     helper.DereferenceString(userGorm.Coordinate),
		Role:           userGorm.Role,
		CreatedAt:      userGorm.CreatedAt,
		UpdatedAt:      userGorm.UpdatedAt,
	}
	return userCore
}
