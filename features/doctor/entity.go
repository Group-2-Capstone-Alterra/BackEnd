package doctor

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/clinic"
)

type Core struct {
	ID             uint
	AdminID        uint
	FullName       string
	Email          string
	Specialization string
	ProfilePicture string
	AvailableDay   clinic.Core
}

type DoctorModel interface {
	AddDoctor(Core) error
	SelectByAdminId(id uint) (*Core, error)
	SelectDoctorById(id uint) (*Core, error)
	SelectAvailDayById(id uint) (*availdaydoctor.Core, error)
	SelectAllDoctor() ([]Core, error)
}

type DoctorService interface {
	AddDoctor(Core) error
}
