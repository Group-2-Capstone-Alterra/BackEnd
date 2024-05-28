package data

import (
	"PetPalApp/features/user"
	"PetPalApp/utils/helper"

	"gorm.io/gorm"
)

type userQuery struct {
	db     *gorm.DB
	helper helper.HelperInterface
}

func New(db *gorm.DB, helper helper.HelperInterface) user.DataInterface {
	return &userQuery{
		db:     db,
		helper: helper,
	}
}

func (u *userQuery) Insert(input user.Core) error {
	userGorm := UserCoreToUserGorm(input, u.helper)
	tx := u.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
