package data

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/availdaydoctor/data"
	"PetPalApp/features/servicedoctor"
	_dataService "PetPalApp/features/servicedoctor/data"
	_serviceData "PetPalApp/features/servicedoctor/data"

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
		AdminID:        doctor.AdminID,
		FullName:       doctor.FullName,
		ProfilePicture: doctor.ProfilePicture,
		About:          doctor.About,
		Price:          doctor.Price,
	}
	tx := dm.db.Create(&doctorGorm)
	if tx.Error != nil {
		return tx.Error
	}

	var doctorData Doctor
	txDoctor := dm.db.Where("admin_id = ?", doctor.AdminID).Find(&doctorData)
	if txDoctor.Error != nil {
		return txDoctor.Error
	}

	var doctorCore = GormToCore(doctorData)

	availdayGorm := data.AvailableDay{
		DoctorID:  doctorCore.ID,
		Monday:    doctor.AvailableDay.Monday,
		Tuesday:   doctor.AvailableDay.Tuesday,
		Wednesday: doctor.AvailableDay.Wednesday,
		Thursday:  doctor.AvailableDay.Thursday,
		Friday:    doctor.AvailableDay.Friday,
	}
	txAvail := dm.db.Create(&availdayGorm)
	if txAvail.Error != nil {
		return txAvail.Error
	}

	serviceGorm := _dataService.ServiceDoctor{
		DoctorID:     doctorCore.ID,
		Vaccinations: doctor.ServiceDoctor.Vaccinations,
		Operations:   doctor.ServiceDoctor.Operations,
		MCU:          doctor.ServiceDoctor.MCU,
	}
	log.Println("[QUERY]serviceGorm", serviceGorm)
	txService := dm.db.Create(&serviceGorm)
	if txService.Error != nil {
		return txService.Error
	}

	return nil
}

func (dm *DoctorModel) SelectByAdminId(id uint) (*doctor.Core, error) {
	var doctorData Doctor
	tx := dm.db.Where("admin_id = ?", id).Find(&doctorData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var doctorCore = GormToCore(doctorData)

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
	log.Println("[Query Doctor - SelectDoctorById] doctorCore", doctorCore)

	return &doctorCore, nil
}

func (dm *DoctorModel) SelectAvailDayById(id uint) (*availdaydoctor.Core, error) {
	var availDay data.AvailableDay

	tx := dm.db.Where("doctor_id = ?", id).Find(&availDay)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var availDayCore = data.AvailGormToCore(availDay)
	return &availDayCore, nil
}

func (dm *DoctorModel) SelectServiceById(id uint) (*servicedoctor.Core, error) {
	var serviceDoctor _serviceData.ServiceDoctor
	log.Println("[Query Doctor - SelectServiceById] iD Param", id)

	tx := dm.db.Where("doctor_id = ?", id).Find(&serviceDoctor)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var serviceDoctorCore = _serviceData.ServiceGormToCore(serviceDoctor)
	return &serviceDoctorCore, nil

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
			About:          v.About,
			Price:          v.Price,
			ProfilePicture: v.ProfilePicture,
		})
	}
	return allDoctorCore, nil
}

func (dm *DoctorModel) PutByIdAdmin(AdminID uint, input doctor.Core) error {

	doctorGorm := Doctor{
		AdminID:        input.AdminID,
		FullName:       input.FullName,
		ProfilePicture: input.ProfilePicture,
		About:          input.About,
		Price:          input.Price,
	}

	tx := dm.db.Model(&Doctor{}).Where("admin_id = ?", AdminID).Updates(&doctorGorm)
	if tx.Error != nil {
		return tx.Error
	}

	var doctorData Doctor
	txDoctor := dm.db.Where("admin_id = ?", AdminID).Find(&doctorData)
	if txDoctor.Error != nil {
		return txDoctor.Error
	}

	var doctorCore = GormToCore(doctorData)

	dm.db.Where("doctor_id = ?", doctorCore.ID).Delete(&data.AvailableDay{})

	availdayGorm := data.AvailableDay{
		DoctorID:  doctorCore.ID,
		Monday:    input.AvailableDay.Monday,
		Tuesday:   input.AvailableDay.Tuesday,
		Wednesday: input.AvailableDay.Wednesday,
		Thursday:  input.AvailableDay.Thursday,
		Friday:    input.AvailableDay.Friday,
	}
	log.Println("availdayGorm", availdayGorm)
	log.Println("doctorCore.ID", doctorCore.ID)
	txAvail := dm.db.Create(&availdayGorm)
	if txAvail.Error != nil {
		return txAvail.Error
	}

	dm.db.Where("doctor_id = ?", doctorCore.ID).Delete(&_dataService.ServiceDoctor{})

	serviceGorm := _dataService.ServiceDoctor{
		DoctorID:            doctorCore.ID,
		Vaccinations:        input.ServiceDoctor.Vaccinations,
		Operations:          input.ServiceDoctor.Operations,
		MCU:                 input.ServiceDoctor.MCU,
		OnlineConsultations: input.ServiceDoctor.OnlineConsultations,
	}
	log.Println("[QUERY]serviceGorm", serviceGorm)
	txService := dm.db.Create(&serviceGorm)
	if txService.Error != nil {
		return txService.Error
	}

	return nil
}

func (dm *DoctorModel) Delete(adminID uint) error {
	tx := dm.db.Where("admin_id = ?", adminID).Delete(&Doctor{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
