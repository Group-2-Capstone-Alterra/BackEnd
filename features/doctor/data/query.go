package data

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/doctor"
	"log"

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
		FullName:       doctor.FullName,
		Email:          doctor.Email,
		Specialization: doctor.Specialization,
	}
	tx := dm.db.Create(&doctorGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dm *DoctorModel) SelectByAdminId(id uint) (*doctor.Core, error) {
	log.Println("[Query Doctor - SelectById]")
	var doctorData Doctor
	log.Println("[Query Doctor - SelectById] AdminID", id)
	tx := dm.db.Where("admin_id = ?", id).Find(&doctorData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var doctorCore = GormToCore(doctorData)
	log.Println("[Query Doctor - SelectById] doctorCore", doctorCore)

	return &doctorCore, nil
}

func (dm *DoctorModel) SelectDoctorById(id uint) (*doctor.Core, error) {
	log.Println("[Query Doctor - SelectById]")
	var doctorData Doctor
	log.Println("[Query Doctor - SelectById] AdminID", id)
	tx := dm.db.Find(&doctorData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var doctorCore = GormToCore(doctorData)
	log.Println("[Query Doctor - SelectById] doctorCore", doctorCore)

	return &doctorCore, nil
}

func (dm *DoctorModel) SelectAvailDayById(id uint) (*availdaydoctor.Core, error) {
	var availDay AvailableDay
	log.Println("[Query Doctor - SelectAvailDayById] iD Param", id)

	tx := dm.db.Where("doctor_id = ?", id).Find(&availDay)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var availDayCore = AvailGormToCore(availDay)
	log.Println("[Query Doctor - SelectAvailDayById] availDayCore", availDayCore)

	return &availDayCore, nil
}

func (dm *DoctorModel) SelectAllDoctor() ([]doctor.Core, error) {
	var allDoctor []Doctor

	tx := dm.db.Find(&allDoctor)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allDoctorCore []doctor.Core
	for _, v := range allDoctor {
		allDoctorCore = append(allDoctorCore, doctor.Core{
			ID:             v.ID,
			AdminID:        v.AdminID,
			FullName:       v.FullName,
			Email:          v.Email,
			Specialization: v.Specialization,
		})
	}
	return allDoctorCore, nil
}
