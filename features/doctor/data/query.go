package data

import (
	"PetPalApp/features/doctor"

	"gorm.io/gorm"
)

type DoctorModel struct {
    db *gorm.DB
}

func New(db *gorm.DB) doctor.DoctorModel {
    return &DoctorModel{
        db: db,
    }
}

func (dm *DoctorModel) AddDoctor(doctor doctor.Core) error {
    doctorGorm := Doctor{
        FullName:      doctor.FullName,
        Email:         doctor.Email,
        Specialization: doctor.Specialization,
    }
    tx := dm.db.Create(&doctorGorm)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}
