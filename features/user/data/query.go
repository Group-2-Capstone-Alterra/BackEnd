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

func (u *userQuery) SelectByEmail(email string) (*user.Core, error) {
	var userData User
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var usercore = UserGormToUserCore(userData)
	return &usercore, nil
}

func (u *userQuery) SelectById(id uint) (*user.Core, error) {
	var userData User
	tx := u.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var usercore = UserGormToUserCore(userData)
	// errAddRedis := UserGormToRedis(u.rdb, userData, ttl)
	// for key, v := range errAddRedis {
	// 	keyWithPrefix := fmt.Sprint(key)
	// 	errIns := u.rdb.Set(u.ctx, keyWithPrefix, v, ttl).Err()
	// 	if errIns != nil {
	// 		log.Print(errIns)
	// 	}
	// }
	return &usercore, nil
}
