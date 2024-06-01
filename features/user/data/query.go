package data

import (
	"PetPalApp/features/user"
	"PetPalApp/utils/helperuser"

	"gorm.io/gorm"
)

type userQuery struct {
	db         *gorm.DB
	helperuser helperuser.HelperuserInterface
}

func New(db *gorm.DB, helperuser helperuser.HelperuserInterface) user.DataInterface {
	return &userQuery{
		db:         db,
		helperuser: helperuser,
	}
}

func (u *userQuery) Insert(input user.Core) error {

	userGorm := UserCoreToUserGorm(input, u.helperuser)
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
	var usercore = UserGormToUserCore(userData, u.helperuser)
	return &usercore, nil
}

func (u *userQuery) SelectById(id uint) (*user.Core, error) {
	var userData User
	tx := u.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var usercore = UserGormToUserCore(userData, u.helperuser)
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

func (u *userQuery) PutById(id uint, input user.Core) error {
	inputGorm := UserCoreToUserGorm(input, u.helperuser)
	tx := u.db.Model(&User{}).Where("id = ?", id).Updates(&inputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *userQuery) Delete(id uint) error {
	tx := u.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
