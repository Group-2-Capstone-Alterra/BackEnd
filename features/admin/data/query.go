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
		Password:       admin.Password,
		NumberPhone:    nil,
		Address:        nil,
		ProfilePicture: nil,
		Coordinate:     nil,
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

	var numberPhone, address, profilePicture string

	if adminData.NumberPhone != nil {
		numberPhone = *adminData.NumberPhone
	}

	if adminData.Address != nil {
		address = *adminData.Address
	}

	if adminData.ProfilePicture != nil {
		profilePicture = *adminData.ProfilePicture
	}

	var admins = admin.Core{
		ID:             adminData.ID,
		FullName:       adminData.FullName,
		Email:          adminData.Email,
		NumberPhone:    numberPhone,
		Address:        address,
		Password:       adminData.Password,
		ProfilePicture: profilePicture,
	}
	return &admins, nil
}

func (am *AdminModel) AdminById(adminid uint) (*admin.Core, error) {
	var adminData Admin
	tx := am.db.First(&adminData, adminid)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var numberPhone, address, profilePicture string

	if adminData.NumberPhone != nil {
		numberPhone = *adminData.NumberPhone
	}

	if adminData.Address != nil {
		address = *adminData.Address
	}

	if adminData.ProfilePicture != nil {
		profilePicture = *adminData.ProfilePicture
	}

	var admin = admin.Core{
		ID:             adminData.ID,
		FullName:       adminData.FullName,
		Email:          adminData.Email,
		NumberPhone:    numberPhone,
		Address:        address,
		ProfilePicture: profilePicture,
		Password:       adminData.Password,
		Coordinate:     *adminData.Coordinate,
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
		if updateData.NumberPhone != "" {
			adminData.NumberPhone = &updateData.NumberPhone
		} else {
			adminData.NumberPhone = nil
		}
	}
	if updateData.Address != "" {
		if updateData.Address != "" {
			adminData.Address = &updateData.Address
		} else {
			adminData.Address = nil
		}
	}
	if updateData.ProfilePicture != "" {
		if updateData.ProfilePicture != "" {
			adminData.ProfilePicture = &updateData.ProfilePicture
		} else {
			adminData.ProfilePicture = nil
		}
	}

	txSave := am.db.Save(&adminData)
	if txSave.Error != nil {
		return txSave.Error
	}

	return nil
}

// func (am *AdminModel) SelectAll(offset uint, sortStr string) ([]admin.AllClinicResponseCore, error) {
// 	var allClinic []AllClinicResponse

// }

func (am *AdminModel) SelectAllAdmin() ([]admin.Core, error) {
	var allAdmin []Admin

	tx := am.db.Find(&allAdmin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allAdminCore []admin.Core
	for _, v := range allAdmin {
		var address, coordinate, profilePicture string

		if v.Address != nil {
			address = *v.Address
		}
		if v.Coordinate != nil {
			coordinate = *v.Coordinate
		}
		if v.ProfilePicture != nil {
			profilePicture = *v.ProfilePicture
		}

		allAdminCore = append(allAdminCore, admin.Core{
			ID:             v.ID,
			FullName:       v.FullName,
			Address:        address,
			Coordinate:     coordinate,
			ProfilePicture: profilePicture,
		})
	}
	return allAdminCore, nil
}
