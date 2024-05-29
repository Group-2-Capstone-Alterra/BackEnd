package data

import (
	"PetPalApp/features/admin"

	"gorm.io/gorm"
)

type AdminModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) admin.AdminModel {
	return &AdminModel{
		db: db,
	}
}

func (am *AdminModel) Register(admin admin.Core) error {
	adminGorm := Admin{
		FullName:       admin.FullName,
		Email:          admin.Email,
		NumberPhone:    admin.NumberPhone,
		Address:        admin.Address,
		Password:       admin.Password,
		ProfilePicture: admin.ProfilePicture,
	}
	tx := am.db.Create(&adminGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (am *AdminModel) AdminByEmail(email string) (*admin.Core, error) {
	var adminData Admin
	tx := am.db.Where("email = ?", email).First(&adminData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var admins = admin.Core{
		ID:             adminData.ID,
		FullName:       adminData.FullName,
		Email:          adminData.Email,
		NumberPhone:    adminData.NumberPhone,
		Address:        adminData.Address,
		Password:       adminData.Password,
		ProfilePicture: adminData.ProfilePicture,
	}
	return &admins, nil
}
