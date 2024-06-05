package doctor

import (
	"PetPalApp/features/availdaydoctor"
	"io"
)

type Core struct {
	ID             uint
	AdminID        uint
	FullName       string
	Email          string
	Specialization string
	ProfilePicture string
	AvailableDay   availdaydoctor.Core
}

type DoctorModel interface {
	AddDoctor(Core) error
	SelectByAdminId(id uint) (*Core, error)
	SelectDoctorById(id uint) (*Core, error)
	SelectAvailDayById(id uint) (*availdaydoctor.Core, error)
	SelectAllDoctor() ([]Core, error)
	PutByIdAdmin(AdminID uint, input Core) error
}

type DoctorService interface {
	AddDoctor(Core) error
	GetDoctorByIdAdmin(adminID uint) (*Core, error)
	GetAvailDoctorByIdDoctor(doctorID uint) (*availdaydoctor.Core, error)
	UpdateByIdAdmin(AdminId uint, input Core, file io.Reader, handlerFilename string) (string, error)
}
