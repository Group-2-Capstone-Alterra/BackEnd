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

func (am *AdminModel) AdminById(adminid uint) (*admin.Core, error) {
	var adminData Admin
	tx := am.db.First(&adminData, adminid)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var admin = admin.Core{
		FullName:       adminData.FullName,
		Email:          adminData.Email,
		NumberPhone:    adminData.NumberPhone,
		Address:        adminData.Address,
		ProfilePicture: adminData.ProfilePicture,
		Coordinate:     adminData.Coordinate,
	}
	return &admin, nil
}

func (am *AdminModel) Delete(adminid uint) error {
	tx := am.db.Delete(&Admin{}, adminid)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (am *AdminModel) Update(adminid uint, updateData admin.Core) error {
	var adminData Admin
	tx := am.db.First(&adminData, adminid)
	if tx.Error != nil {
		return tx.Error
	}

	if updateData.FullName != "" {
		adminData.FullName = updateData.FullName
	}
	if updateData.Email != "" {
		adminData.Email = updateData.Email
	}
	if updateData.NumberPhone != "" {
		adminData.NumberPhone = updateData.NumberPhone
	}
	if updateData.Address != "" {
		adminData.Address = updateData.Address
	}
	if updateData.ProfilePicture != "" {
		adminData.ProfilePicture = updateData.ProfilePicture
	}

	txSave := am.db.Save(&adminData)
	if txSave.Error != nil {
		return txSave.Error
	}

	return nil
}
